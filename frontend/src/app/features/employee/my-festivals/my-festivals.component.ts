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
    this.festivalService.getMyFestivalsEmployee().subscribe({
      next: (response) => {
        this.festivals = response;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error fetching festivals:', error);
        this.snackbarService.show('Error fetching festivals');
        this.isLoading = true;
      },
    });
  }

  onViewClick(festival: Festival): void {
    this.router.navigate(['employee/my-festivals', festival.id]);
  }
}
