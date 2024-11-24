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
  ConfirmationDialog,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { EditFestivalComponent } from '../edit-festival/edit-festival.component';
import { animate, style, transition, trigger } from '@angular/animations';
import { MatMenuModule } from '@angular/material/menu';

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
  onViewEmployeesClick(arg0: Festival) {
    throw new Error('Method not implemented.');
  }
  festival: Festival | null = null;
  isLoading: boolean = true;
  currentImageIndex: number = 0;
  previousImageIndex: number = 0;
  employeeCount: number = 0;

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadEmployeeCount();
  }

  goBack() {
    this.router.navigate(['organizer/my-festivals']);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          console.log('Festival: ', festival);
          this.festival = festival;
          this.isLoading = false;
          this.currentImageIndex = 0;
          this.previousImageIndex = 0;
        },
        error: (error) => {
          console.log('Error fetching festival information: ', error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
          this.isLoading = true;
        },
      });
    }
  }

  loadEmployeeCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getEmployeeCount(Number(id)).subscribe({
        next: (count) => {
          console.log('Employee count: ', count);
          this.employeeCount = count;
        },
        error: (error) => {
          console.log('Error fetching employee count: ', error);
          this.snackbarService.show('Error getting employee count');
          this.employeeCount = 0;
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
    console.log('Delete clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
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
    console.log('Edit clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(EditFestivalComponent, {
      data: {
        id: this.festival?.id,
        name: this.festival?.name,
        description: this.festival?.description,
        startDate: this.festival?.startDate,
        endDate: this.festival?.endDate,
        capacity: this.festival?.capacity,
        address: this.festival?.address,
      },
      width: '800px',
      height: '535px',
    });

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        console.log('Festival updated successfully');
        this.loadFestival();
      }
    });
  }

  onPublishClick(): void {
    console.log('Publish clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Publish Festival',
        message: `Are you sure you want to publish ${this.festival?.name}?`,
        confirmButtonText: 'Publish',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Publishing festival:', this.festival?.name);
        this.festivalService
          .publishFestival(Number(this.festival?.id))
          .subscribe({
            next: () => {
              console.log('Festival published:', this.festival?.name);
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
    console.log('Cancel clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Cancel Festival',
        message: `Are you sure you want to cancel ${this.festival?.name}?`,
        confirmButtonText: 'Cancel',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Cancelling festival:', this.festival?.name);
        this.festivalService
          .cancelFestival(Number(this.festival?.id))
          .subscribe({
            next: () => {
              console.log('Festival cancelled:', this.festival?.name);
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
    console.log('Complete clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Complete Festival',
        message: `Are you sure you want to complete ${this.festival?.name}?`,
        confirmButtonText: 'Complete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Completing festival:', this.festival?.name);
        this.festivalService
          .completeFestival(Number(this.festival?.id))
          .subscribe({
            next: () => {
              console.log('Festival completed:', this.festival?.name);
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
    console.log('Store open clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Open Store',
        message: `Are you sure you want to open store for ${this.festival?.name}?`,
        confirmButtonText: 'Open',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Opening store for festival:', this.festival?.name);
        this.festivalService
          .openFestivalStore(Number(this.festival?.id))
          .subscribe({
            next: () => {
              console.log('Store opened for festival:', this.festival?.name);
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
    console.log('Store close clicked for:', this.festival?.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Close Store',
        message: `Are you sure you want to close store for ${this.festival?.name}?`,
        confirmButtonText: 'Close',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Closing store for festival:', this.festival?.name);
        this.festivalService
          .closeFestivalStore(Number(this.festival?.id))
          .subscribe({
            next: () => {
              console.log('Store closed for festival:', this.festival?.name);
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
}
