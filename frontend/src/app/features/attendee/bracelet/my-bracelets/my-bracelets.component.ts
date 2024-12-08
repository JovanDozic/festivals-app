import { Component, inject, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { OrderService } from '../../../../services/festival/order.service';
import { OrderPreviewDTO } from '../../../../models/festival/festival.model';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-my-bracelets',
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
  templateUrl: './my-bracelets.component.html',
  styleUrls: [
    './my-bracelets.component.scss',
    '../../../../app.component.scss',
  ],
})
export class MyBraceletsComponent {
  private router = inject(Router);
  private snackbarService = inject(SnackbarService);
  private orderService = inject(OrderService);

  isLoading: boolean = true;
  bracelets: any[] = [
    {
      name: 'jovans',
      status: 'Activated',
    },
    {
      name: 'navojs',
      status: 'Activated',
    },
    {
      name: 'tomorrowland',
      status: 'Activated',
    },
    {
      name: 'losercxl',
      status: 'Activated',
    },
    {
      name: 'oldedwo',
      status: 'Activated',
    },
    {
      name: 'loovfgerel',
      status: 'Activated',
    },
  ];

  filterOptions: string[] = [
    'All',
    'Pending',
    'Issued',
    'Activated',
    'Help Requested', // ovo mozda i ne stavljamo ovde jer je tesko to dobaviti u orders
    'Rejected',
  ];
  selectedChip: string = 'All'; // todo: change to activated

  constructor() {}

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
    this.isLoading = false;
    // throw new Error('Method not implemented.');
    // this.orderService.getMyOrders().subscribe({
    //   next: (orders) => {
    //     this.orders = orders;
    //     this.isLoading = false;
    //   },
    //   error: (error) => {
    //     console.log(error);
    //     this.snackbarService.show('Failed to load orders');
    //   },
    // });
  }

  // get filteredBracelets(): OrderPreviewDTO[] {
  //   if (!this.orders || this.orders.length === 0) {
  //     return [];
  //   }
  //   if (this.selectedChip === 'Upcoming Festivals') {
  //     return this.orders.filter(
  //       (order) => new Date(order.festival.startDate) > new Date(),
  //     );
  //   } else if (this.selectedChip === 'Past Festivals') {
  //     return this.orders.filter(
  //       (order) => new Date(order.festival.startDate) < new Date(),
  //     );
  //   } else {
  //     return this.orders;
  //   }
  // }

  onViewClick() {
    // this.router.navigate(['my-orders', order.orderId]);
  }
}
