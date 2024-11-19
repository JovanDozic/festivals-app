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
}
