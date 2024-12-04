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
import { OrderDTO } from '../../../../models/festival/festival.model';

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

  order: OrderDTO | null = null;

  constructor() {}

  ngOnInit() {
    this.loadOrder();
  }

  goBack() {
    this.router.navigate([`my-orders/`]);
  }

  loadOrder() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getOrder(parseInt(id)).subscribe({
        next: (order) => {
          console.log('Order: ', order);
          this.order = order;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.isLoading = false;
          this.snackbarService.show('Error loading order');
        },
      });
    }
  }
}
