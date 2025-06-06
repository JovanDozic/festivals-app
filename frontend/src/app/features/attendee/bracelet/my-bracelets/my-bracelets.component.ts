import { Component, inject, OnInit } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
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
import { OrderDTO } from '../../../../models/festival/festival.model';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';
import { FormsModule } from '@angular/forms';
import { ActivateBraceletComponent } from '../activate-bracelet/activate-bracelet.component';
import { TopUpBraceletComponent } from '../top-up-bracelet/top-up-bracelet.component';

@Component({
  selector: 'app-my-bracelets',
  imports: [
    CommonModule,
    RouterLink,
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
  templateUrl: './my-bracelets.component.html',
  styleUrls: [
    './my-bracelets.component.scss',
    '../../../../app.component.scss',
  ],
})
export class MyBraceletsComponent implements OnInit {
  private router = inject(Router);
  private snackbarService = inject(SnackbarService);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

  isLoading = true;
  orders: OrderDTO[] = [];

  filterOptions: string[] = [
    'All',
    'Not Issued',
    'Issued and Shipped',
    'Active',
    'Deactivated',
    'Help Requested',
    'Rejected',
  ];
  selectedChip = 'Active';

  ngOnInit() {
    this.loadBracelets();
  }

  private getColorFromName(name: string) {
    let hash = 0;
    for (let i = 0; i < name.length; i++) {
      hash = name.charCodeAt(i) + ((hash << 5) - hash);
    }
    const r = (hash >> 16) & 0xff;
    const g = (hash >> 8) & 0xff;
    const b = hash & 0xff;
    return { r, g, b };
  }

  getBraceletStyle(name: string) {
    const { r, g, b } = this.getColorFromName(name);
    const rgba1 = `rgba(${r},${g},${b},0.2)`;
    const rgba2 = `rgba(${(r + 128) % 256},${(g + 128) % 256},${(b + 128) % 256},0.2)`;
    return {
      '--random-color-1': rgba1,
      '--random-color-2': rgba2,
    };
  }

  loadBracelets() {
    this.orderService.getMyBracelets().subscribe({
      next: (orders) => {
        this.orders = orders;
        this.isLoading = false;
      },
      error: (error) => {
        console.log(error);
        this.snackbarService.show('Failed to load bracelets');
      },
    });
  }

  get filteredOrders(): OrderDTO[] {
    if (!this.orders || this.orders.length === 0) {
      return [];
    }
    if (this.selectedChip === 'Not Issued') {
      return this.orders.filter((order) => !order.braceletStatus);
    } else if (this.selectedChip === 'Issued and Shipped') {
      return this.orders.filter((order) => order.braceletStatus === 'ISSUED');
    } else if (this.selectedChip === 'Active') {
      return this.orders.filter(
        (order) => order.braceletStatus === 'ACTIVATED',
      );
    } else if (this.selectedChip === 'Deactivated') {
      return this.orders.filter(
        (order) => order.braceletStatus === 'DEACTIVATED',
      );
    } else if (this.selectedChip === 'Help Requested') {
      return this.orders.filter(
        (order) => order.braceletStatus === 'HELP_REQUESTED',
      );
    } else if (this.selectedChip === 'Rejected') {
      return this.orders.filter((order) => order.braceletStatus === 'REJECTED');
    } else {
      return this.orders;
    }
  }

  onViewClick() {
    // this.router.navigate(['my-orders', order.orderId]);
  }

  getBraceletStatusText(status: string): string {
    switch (status) {
      case 'ISSUED':
        return 'Issued and Shipped';
      case 'ACTIVATED':
        return 'Activated';
      case 'HELP_REQUESTED':
        return 'Help Requested';
      case 'DEACTIVATED':
        return 'Deactivated';
      case 'REJECTED':
        return 'Rejected';
      default:
        return 'Not Issued';
    }
  }

  activateBracelet(order: OrderDTO) {
    const dialogRef = this.dialog.open(ActivateBraceletComponent, {
      data: {
        order: order,
      },
      width: '800px',
      height: '425px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadBracelets();
      }
    });
  }

  topUpBracelet(order: OrderDTO) {
    const dialogRef = this.dialog.open(TopUpBraceletComponent, {
      data: {
        order: order,
      },
      width: '800px',
      height: '325px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadBracelets();
      }
    });
  }
}
