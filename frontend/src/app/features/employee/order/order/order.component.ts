import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../../services/festival/festival.service';
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { ItemService } from '../../../../services/festival/item.service';
import { UserService } from '../../../../services/user/user.service';
import { OrderService } from '../../../../services/festival/order.service';
import {
  Festival,
  OrderDTO,
  OrderPreviewDTO,
} from '../../../../models/festival/festival.model';
import { AddressResponse } from '../../../../models/common/address-response.model';
import { UserProfileResponse } from '../../../../models/user/user-profile-response.model';
import { IssueBraceletComponent } from '../issue-bracelet/issue-bracelet.component';

@Component({
  selector: 'app-order',
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
  ],
  templateUrl: './order.component.html',
  styleUrls: ['./order.component.scss', '../../../../app.component.scss'],
})
export class OrderComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private itemService = inject(ItemService);
  private userService = inject(UserService);
  private snackbarService = inject(SnackbarService);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

  isLoading: boolean = true;

  festival: Festival | null = null;

  order: OrderDTO | null = null;
  userProfile: UserProfileResponse | null = null;
  address: AddressResponse | null = null;

  constructor() {}

  ngOnInit() {
    this.loadFestival();

    this.loadOrder();
    this.loadUser();
  }

  goBack() {
    this.router.navigate([`employee/my-festivals/${this.festival?.id}/orders`]);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
          this.isLoading = true;
        },
      });
    }
  }

  loadOrder() {
    const id = this.route.snapshot.paramMap.get('orderId');
    if (id) {
      this.itemService.getOrder(parseInt(id)).subscribe({
        next: (order) => {
          this.order = order;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.isLoading = false;
          if (error.status === 404) {
            this.snackbarService.show('Order not found');
          } else {
            this.snackbarService.show('Error loading order');
          }
        },
      });
    }
  }

  loadUser() {
    this.userService.getUserProfile().subscribe({
      next: (userProfile) => {
        this.userProfile = userProfile;
        if (userProfile.address) this.address = userProfile.address;
        this.isLoading = false;
      },
      error: (error) => {
        console.log(error);
        this.snackbarService.show('Error getting user information');
        this.userProfile = null;
        this.isLoading = true;
      },
    });
  }

  onIssueBraceletClick() {
    if (this.order) {
      const dialogRef = this.dialog.open(IssueBraceletComponent, {
        data: {
          festivalId: this.festival?.id,
          orderId: this.order.orderId,
          festivalTicketId: this.order.festivalTicketId,
          attendeeUsername: this.order.attendee.username,
        },
        width: '800px',
        height: 'auto',
        disableClose: true,
      });

      dialogRef.afterClosed().subscribe((result) => {
        if (result) {
          this.snackbarService.show('Bracelet issued successfully!');
          this.loadOrder();
        }
      });
    }
  }

  onPrintShippingLabelClick() {
    if (this.order && this.order.orderId) {
      this.snackbarService.show('Printing shipping label...');
      this.orderService.printShippingLabel(this.order.orderId).subscribe({
        next: (blob) => {
          const url = URL.createObjectURL(blob);
          window.open(url, '_blank');
          this.snackbarService.show('Shipping label opened successfully');
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error printing shipping label');
        },
      });
    }
  }

  reviewHelpRequest(order: OrderDTO) {}
}
