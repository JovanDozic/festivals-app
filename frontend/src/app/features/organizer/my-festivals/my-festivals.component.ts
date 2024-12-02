import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { Festival } from '../../../models/festival/festival.model';
import { FestivalService } from '../../../services/festival/festival.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';
import { MatDialog } from '@angular/material/dialog';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { Router, RouterModule } from '@angular/router';
import { EditFestivalComponent } from '../edit-festival/edit-festival.component';

@Component({
  selector: 'app-my-festivals',
  templateUrl: './my-festivals.component.html',
  styleUrls: ['./my-festivals.component.scss', '../../../app.component.scss'],
  imports: [
    CommonModule,
    RouterModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatMenuModule,
    NgxSkeletonLoaderModule,
  ],
})
export class MyFestivalsComponent implements OnInit {
  festivals: Festival[] = [];
  isLoading = true;

  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);
  private router = inject(Router);

  ngOnInit(): void {
    this.loadFestivals();
  }

  getSkeletonBgColor(): string {
    const isDarkTheme =
      document.documentElement.getAttribute('data-theme') === 'dark';
    return isDarkTheme ? '#494d8aaa' : '#e0e0ff';
  }

  loadFestivals(): void {
    this.festivalService.getMyFestivals().subscribe({
      next: (response) => {
        // setTimeout(() => {
        this.festivals = response;
        this.isLoading = false;
        // }, 250);
      },
      error: (error) => {
        console.error('Error fetching festivals:', error);
        this.snackbarService.show('Error fetching festivals');
        this.isLoading = true;
      },
    });
  }

  onDeleteClick(festival: Festival): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Delete Festival',
        message: `Are you sure you want to delete ${festival.name}?`,
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService.deleteFestival(festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Festival deleted');
            this.loadFestivals();
          },
          error: (error) => {
            console.error('Error deleting festival:', error);
            this.snackbarService.show('Error deleting festival');
          },
        });
      }
    });
  }

  onEditClick(festival: Festival): void {
    const dialogRef = this.dialog.open(EditFestivalComponent, {
      data: {
        id: festival.id,
        name: festival.name,
        description: festival.description,
        startDate: festival.startDate,
        endDate: festival.endDate,
        capacity: festival.capacity,
        address: festival.address,
      },
      width: '800px',
      height: '535px',
    });

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        this.loadFestivals();
      }
    });
  }

  onViewClick(festival: Festival): void {
    this.router.navigate(['organizer/my-festivals', festival.id]);
  }

  onPublishClick(festival: Festival): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Publish Festival',
        message: `Are you sure you want to publish ${festival.name}?`,
        confirmButtonText: 'Publish',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService.publishFestival(festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Festival published');
            this.loadFestivals();
          },
          error: (error) => {
            console.error('Error published festival:', error);
            this.snackbarService.show('Error publish festival');
          },
        });
      }
    });
  }

  onCancelClick(festival: Festival): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Cancel Festival',
        message: `Are you sure you want to cancel ${festival.name}?`,
        confirmButtonText: 'Cancel',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService.cancelFestival(festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Festival cancelled');
            this.loadFestivals();
          },
          error: (error) => {
            console.error('Error cancelling festival:', error);
            this.snackbarService.show('Error cancelling festival');
          },
        });
      }
    });
  }

  onCompleteClick(festival: Festival): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Complete Festival',
        message: `Are you sure you want to complete ${festival.name}?`,
        confirmButtonText: 'Complete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService.completeFestival(festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Festival completed');
            this.loadFestivals();
          },
          error: (error) => {
            console.error('Error completing festival:', error);
            this.snackbarService.show('Error completing festival');
          },
        });
      }
    });
  }

  onStoreOpenClick(festival: Festival): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Open Store',
        message: `Are you sure you want to open store for ${festival.name}?`,
        confirmButtonText: 'Open',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService.openFestivalStore(festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Store opened');
            this.loadFestivals();
          },
          error: (error) => {
            console.error('Error opening store:', error);
            this.snackbarService.show('Error opening store');
          },
        });
      }
    });
  }

  onStoreCloseClick(festival: Festival): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Close Store',
        message: `Are you sure you want to close store for ${festival.name}?`,
        confirmButtonText: 'Close',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.festivalService.closeFestivalStore(festival.id).subscribe({
          next: () => {
            this.snackbarService.show('Store closed');
            this.loadFestivals();
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
