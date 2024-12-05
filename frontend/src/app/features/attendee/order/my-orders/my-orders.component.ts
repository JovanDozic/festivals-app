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
import { OrderPreviewDTO } from '../../../../models/festival/festival.model';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-all-orders',
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
  templateUrl: './my-orders.component.html',
  styleUrls: ['./my-orders.component.scss', '../../../../app.component.scss'],
})
export class MyOrdersComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private itemService = inject(ItemService);
  private userService = inject(UserService);
  private snackbarService = inject(SnackbarService);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

  isLoading: boolean = true;
  orders: OrderPreviewDTO[] = [];

  filterOptions: string[] = ['All', 'Upcoming Festivals', 'Past Festivals'];
  selectedChip: string = 'All';

  constructor() {}

  getSkeletonBgColor(): string {
    const isDarkTheme =
      document.documentElement.getAttribute('data-theme') === 'dark';
    return isDarkTheme ? '#494d8aaa' : '#e0e0ff';
  }

  ngOnInit() {
    this.loadOrders();
  }

  loadOrders() {
    this.orderService.getMyOrders().subscribe({
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

  get filteredOrders(): OrderPreviewDTO[] {
    if (this.selectedChip === 'Upcoming Festivals') {
      return this.orders.filter(
        (order) => new Date(order.festival.startDate) > new Date(),
      );
    } else if (this.selectedChip === 'Past Festivals') {
      return this.orders.filter(
        (order) => new Date(order.festival.startDate) < new Date(),
      );
    } else {
      return this.orders;
    }
  }

  onViewClick(order: OrderPreviewDTO) {
    this.router.navigate(['my-orders', order.orderId]);
  }
}
