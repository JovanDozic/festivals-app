import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { OrderService } from '../../../../services/festival/order.service';
import {
  Festival,
  OrderPreviewDTO,
} from '../../../../models/festival/festival.model';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';
import { FormsModule } from '@angular/forms';
import { FestivalService } from '../../../../services/festival/festival.service';
import { IssueBraceletComponent } from '../issue-bracelet/issue-bracelet.component';
import { ReviewHelpRequestComponent } from '../review-help-request/review-help-request.component';

@Component({
  selector: 'app-festival-orders',
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
    NgxSkeletonLoaderModule,
  ],
  templateUrl: './festival-orders.component.html',
  styleUrls: [
    './festival-orders.component.scss',
    '../../../../app.component.scss',
  ],
})
export class FestivalOrdersComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private snackbarService = inject(SnackbarService);
  private orderService = inject(OrderService);
  private festivalService = inject(FestivalService);
  private dialog = inject(MatDialog);

  festival: Festival | null = null;
  isLoading = true;
  orders: OrderPreviewDTO[] = [];

  filterOptions: string[] = [
    'All',
    'Pending',
    'Issued',
    'Activated',
    'Help Requested',
    'Rejected',
  ];
  selectedChip = 'All';

  getSkeletonBgColor(): string {
    const isDarkTheme =
      document.documentElement.getAttribute('data-theme') === 'dark';
    return isDarkTheme ? '#494d8aaa' : '#e0e0ff';
  }

  ngOnInit() {
    this.loadFestival();
    this.loadOrders();
  }

  goBack() {
    this.router.navigate(['employee/my-festivals/', this.festival?.id]);
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

  loadOrders() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.orderService.getFestivalOrders(Number(id)).subscribe({
        next: (orders) => {
          this.orders = orders;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Failed to load orders');
        },
      });
    }
  }

  get filteredOrders(): OrderPreviewDTO[] {
    if (!this.orders || this.orders.length === 0) {
      return [];
    }
    if (this.selectedChip.includes('Pending')) {
      return this.orders.filter((order) => !order.braceletStatus);
    } else if (this.selectedChip.includes('Issued')) {
      return this.orders.filter((order) => order.braceletStatus === 'ISSUED');
    } else if (this.selectedChip.includes('Activated')) {
      return this.orders.filter(
        (order) => order.braceletStatus === 'ACTIVATED',
      );
    } else if (this.selectedChip.includes('Help Requested')) {
      return this.orders.filter(
        (order) => order.braceletStatus === 'HELP_REQUESTED',
      );
    } else if (this.selectedChip.includes('Rejected')) {
      return this.orders.filter((order) => order.braceletStatus === 'REJECTED');
    } else {
      return this.orders;
    }
  }

  onViewClick(order: OrderPreviewDTO) {
    this.router.navigate([
      `employee/my-festivals/${this.festival?.id}/orders/${order.orderId}`,
    ]);
  }

  onIssueBraceletClick(order: OrderPreviewDTO) {
    const dialogRef = this.dialog.open(IssueBraceletComponent, {
      data: {
        festivalId: this.festival?.id,
        orderId: order.orderId,
        festivalTicketId: order.festivalTicketId,
        attendeeUsername: order.username,
      },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.snackbarService.show('Bracelet issued successfully');
        this.loadOrders();
      }
    });
  }

  reviewHelpRequest(order: OrderPreviewDTO) {
    const dialogRef = this.dialog.open(ReviewHelpRequestComponent, {
      data: {
        festivalId: this.festival?.id,
        orderId: order.orderId,
        festivalTicketId: order.festivalTicketId,
        attendeeUsername: order.username,
        braceletId: order.braceletId,
      },
      width: '800px',
      height: '715px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.snackbarService.show(
          'Activation help request reviewed successfully',
        );
        this.loadOrders();
      }
    });
  }
}
