import { Component, inject, OnInit } from '@angular/core';
import { Festival } from '../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../services/festival/festival.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { EditFestivalComponent } from '../edit-festival/edit-festival.component';
import { animate, style, transition, trigger } from '@angular/animations';
import { MatMenuModule } from '@angular/material/menu';
import { ItemService } from '../../../services/festival/item.service';
import { OrderService } from '../../../services/festival/order.service';

@Component({
  selector: 'app-festival',
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
  ],
  templateUrl: './festival.component.html',
  styleUrls: ['./festival.component.scss', '../../../app.component.scss'],
  animations: [
    trigger('fadeAnimation', [
      transition(':increment', [
        style({ opacity: 0 }),
        animate('500ms ease-in', style({ opacity: 1 })),
      ]),
      transition(':decrement', [
        style({ opacity: 0 }),
        animate('500ms ease-in', style({ opacity: 1 })),
      ]),
    ]),
  ],
})
export class FestivalComponent implements OnInit {
  festival: Festival | null = null;
  isLoading = true;
  currentImageIndex = 0;
  previousImageIndex = 0;
  employeesCount = 0;
  ticketTypesCount = 0;
  packageAddonsCount = 0;
  ordersCount = 0;

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private itemService = inject(ItemService);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadEmployeesCount();
    this.loadTicketTypesCount();
    this.loadPackageAddonsCount();
    this.loadOrdersCount();
  }

  goBack() {
    this.router.navigate(['organizer/my-festivals']);
  }

  onViewEmployeesClick() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/employees`,
    ]);
  }

  onViewTicketTypesClick() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/ticket-types`,
    ]);
  }

  onViewPackageAddonsClick() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/package-addons`,
    ]);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
          this.isLoading = false;
          this.currentImageIndex = 0;
          this.previousImageIndex = 0;
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

  loadEmployeesCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getEmployeeCount(Number(id)).subscribe({
        next: (count) => {
          this.employeesCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting employee count');
          this.employeesCount = 0;
        },
      });
    }
  }

  loadTicketTypesCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getTicketTypesCount(Number(id)).subscribe({
        next: (count) => {
          this.ticketTypesCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting ticket types count');
          this.ticketTypesCount = 0;
        },
      });
    }
  }

  loadPackageAddonsCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getAllPackageAddonsCount(Number(id)).subscribe({
        next: (count) => {
          this.packageAddonsCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting package addons count');
          this.packageAddonsCount = 0;
        },
      });
    }
  }

  loadOrdersCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.orderService.getOrdersCount(Number(id)).subscribe({
        next: (count) => {
          this.ordersCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting orders count');
          this.ordersCount = 0;
        },
      });
    }
  }

  nextImage() {
    if (this.festival && this.festival.images) {
      this.previousImageIndex = this.currentImageIndex;
      this.currentImageIndex =
        (this.currentImageIndex + 1) % this.festival.images.length;
    }
  }

  previousImage() {
    if (this.festival && this.festival.images) {
      this.previousImageIndex = this.currentImageIndex;
      this.currentImageIndex =
        (this.currentImageIndex - 1 + this.festival.images.length) %
        this.festival.images.length;
    }
  }

  onDeleteClick(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Delete Festival',
        message: `Are you sure you want to delete ${this.festival?.name}?`,
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm && this.festival) {
        this.festivalService.deleteFestival(this.festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Festival deleted');
            this.router.navigate(['organizer/my-festivals']);
          },
          error: (error) => {
            console.error('Error deleting festival:', error);
            this.snackbarService.show('Error deleting festival');
          },
        });
      }
    });
  }

  onEditClick(): void {
    const dialogRef = this.dialog.open(EditFestivalComponent, {
      data: this.festival,
      width: '800px',
      height: '535px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        this.loadFestival();
      }
    });
  }

  onPublishClick(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Publish Festival',
        message: `Are you sure you want to publish ${this.festival?.name}?`,
        confirmButtonText: 'Publish',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService
          .publishFestival(Number(this.festival?.id))
          .subscribe({
            next: () => {
              this.snackbarService.show('Festival published');
              this.loadFestival();
            },
            error: (error) => {
              console.error('Error published festival:', error);
              this.snackbarService.show('Error publish festival');
            },
          });
      }
    });
  }

  onCancelClick(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Cancel Festival',
        message: `Are you sure you want to cancel ${this.festival?.name}?`,
        confirmButtonText: 'Cancel',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService
          .cancelFestival(Number(this.festival?.id))
          .subscribe({
            next: () => {
              this.snackbarService.show('Festival cancelled');
              this.loadFestival();
            },
            error: (error) => {
              console.error('Error cancelling festival:', error);
              this.snackbarService.show('Error cancelling festival');
            },
          });
      }
    });
  }

  onCompleteClick(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Complete Festival',
        message: `Are you sure you want to complete ${this.festival?.name}?`,
        confirmButtonText: 'Complete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService
          .completeFestival(Number(this.festival?.id))
          .subscribe({
            next: () => {
              this.snackbarService.show('Festival completed');
              this.loadFestival();
            },
            error: (error) => {
              console.error('Error completing festival:', error);
              this.snackbarService.show('Error completing festival');
            },
          });
      }
    });
  }

  onStoreOpenClick(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Open Store',
        message: `Are you sure you want to open store for ${this.festival?.name}?`,
        confirmButtonText: 'Open',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService
          .openFestivalStore(Number(this.festival?.id))
          .subscribe({
            next: () => {
              this.snackbarService.show('Store opened');
              this.loadFestival();
            },
            error: (error) => {
              console.error('Error opening store:', error);
              this.snackbarService.show('Error opening store');
            },
          });
      }
    });
  }

  onStoreCloseClick(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Close Store',
        message: `Are you sure you want to close store for ${this.festival?.name}?`,
        confirmButtonText: 'Close',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService
          .closeFestivalStore(Number(this.festival?.id))
          .subscribe({
            next: () => {
              this.snackbarService.show('Store closed');
              this.loadFestival();
            },
            error: (error) => {
              console.error('Error closing store:', error);
              this.snackbarService.show('Error closing store');
            },
          });
      }
    });
  }

  onViewOrders() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/orders`,
    ]);
  }
}
