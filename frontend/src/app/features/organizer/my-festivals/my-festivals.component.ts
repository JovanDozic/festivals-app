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
  ConfirmationDialog,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';

@Component({
  selector: 'app-my-festivals',
  templateUrl: './my-festivals.component.html',
  styleUrls: ['./my-festivals.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatMenuModule,
    NgxSkeletonLoaderModule,
  ],
})
export class MyFestivalsComponent implements OnInit {
  onUnpublishClick(_t13: Festival) {
    throw new Error('Method not implemented.');
  }
  festivals: Festival[] = [];
  isLoading: boolean = true; // Loading state

  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);

  ngOnInit(): void {
    this.loadFestivals();
  }

  loadFestivals(): void {
    this.festivalService.getMyFestivals().subscribe({
      next: (response) => {
        console.log('Festivals:', response);
        this.festivals = response;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error fetching festivals:', error);
        this.snackbarService.show('Error fetching festivals');
        // this.isLoading = false;
      },
    });
  }

  onDeleteClick(festival: Festival): void {
    console.log('Delete clicked for:', festival.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Delete Festival',
        message: `Are you sure you want to delete ${festival.name}?`,
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Deleting festival:', festival.name);
        this.festivalService.deleteFestival(festival.id).subscribe({
          next: () => {
            console.log('Festival deleted:', festival.name);
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
    console.log('Edit clicked for:', festival.name);
  }

  onViewClick(festival: Festival): void {
    console.log('View clicked for:', festival.name);
  }

  onPublishClick(festival: Festival): void {
    console.log('Publish clicked for:', festival.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Publish Festival',
        message: `Are you sure you want to publish ${festival.name}?`,
        confirmButtonText: 'Publish',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Publishing festival:', festival.name);
        this.festivalService.publishFestival(festival.id).subscribe({
          next: () => {
            console.log('Festival published:', festival.name);
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
    console.log('Cancel clicked for:', festival.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Cancel Festival',
        message: `Are you sure you want to cancel ${festival.name}?`,
        confirmButtonText: 'Cancel',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Cancelling festival:', festival.name);
        this.festivalService.cancelFestival(festival.id).subscribe({
          next: () => {
            console.log('Festival cancelled:', festival.name);
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
    console.log('Complete clicked for:', festival.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Complete Festival',
        message: `Are you sure you want to complete ${festival.name}?`,
        confirmButtonText: 'Complete',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Completing festival:', festival.name);
        this.festivalService.completeFestival(festival.id).subscribe({
          next: () => {
            console.log('Festival completed:', festival.name);
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
    console.log('Store open clicked for:', festival.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Open Store',
        message: `Are you sure you want to open store for ${festival.name}?`,
        confirmButtonText: 'Open',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Opening store for festival:', festival.name);
        this.festivalService.openFestivalStore(festival.id).subscribe({
          next: () => {
            console.log('Store opened for festival:', festival.name);
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
    console.log('Store close clicked for:', festival.name);
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Close Store',
        message: `Are you sure you want to close store for ${festival.name}?`,
        confirmButtonText: 'Close',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        console.log('Closing store for festival:', festival.name);
        this.festivalService.closeFestivalStore(festival.id).subscribe({
          next: () => {
            console.log('Store closed for festival:', festival.name);
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
