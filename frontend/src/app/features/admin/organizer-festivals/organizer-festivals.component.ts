import { Component, inject, OnInit } from '@angular/core';
import { Employee, Festival } from '../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../services/festival/festival.service';
import { SnackbarService } from '../../../services/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { MatTableModule } from '@angular/material/table';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { UserService } from '../../../services/user/user.service';
import { UserProfileResponse } from '../../../models/user/user-responses';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';

@Component({
  selector: 'app-organizer-festivals',
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
    MatTableModule,
    NgxSkeletonLoaderModule,
  ],
  templateUrl: './organizer-festivals.component.html',
  styleUrls: [
    './organizer-festivals.component.scss',
    '../../../app.component.scss',
  ],
})
export class OrganizerFestivalsComponent {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);
  private userService = inject(UserService);

  festivals: Festival[] = [];
  isLoading = true;
  organizer: UserProfileResponse | undefined;
  displayedColumns = ['username', 'email', 'name', 'phoneNumber', 'actions'];

  organizerId: number = 0;

  ngOnInit() {
    this.loadOrganizer();
    this.loadFestivals();
  }

  goBack() {
    this.router.navigate([`admin/users/${this.organizerId}`]);
  }

  getSkeletonBgColor(): string {
    const isDarkTheme =
      document.documentElement.getAttribute('data-theme') === 'dark';
    return isDarkTheme ? '#494d8aaa' : '#e0e0ff';
  }

  loadFestivals(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestivalsByOrganizerId(Number(id)).subscribe({
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
  }

  loadOrganizer() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.organizerId = Number(id);
      this.userService.getUserById(Number(id)).subscribe((response) => {
        this.organizer = response;
      });
    }
  }
}
