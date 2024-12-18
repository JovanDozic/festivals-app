import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { Festival } from '../../../models/festival/festival.model';
import { FestivalService } from '../../../services/festival/festival.service';
import { SnackbarService } from '../../../services/snackbar/snackbar.service';
import { NgxSkeletonLoaderModule } from 'ngx-skeleton-loader';
import { MatDialog } from '@angular/material/dialog';
import { Router, RouterModule } from '@angular/router';
import { MatChipsModule } from '@angular/material/chips';
import { FormsModule } from '@angular/forms';
import { FestivalComponent } from '../../organizer/festival/festival.component';

@Component({
  selector: 'app-all-festivals',
  imports: [
    CommonModule,
    FormsModule,
    RouterModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatMenuModule,
    NgxSkeletonLoaderModule,
    MatChipsModule,
  ],
  templateUrl: './all-festivals.component.html',
  styleUrls: ['./all-festivals.component.scss', '../../../app.component.scss'],
})
export class AllFestivalsComponent implements OnInit {
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  // private dialog = inject(MatDialog);
  private router = inject(Router);

  festivals: Festival[] = [];
  isLoading = true;

  filterOptions: string[] = ['All', 'Upcoming', 'Past'];
  selectedChip = 'All';

  ngOnInit(): void {
    this.loadFestivals();
  }

  getSkeletonBgColor(): string {
    const isDarkTheme =
      document.documentElement.getAttribute('data-theme') === 'dark';
    return isDarkTheme ? '#494d8aaa' : '#e0e0ff';
  }

  loadFestivals(): void {
    this.festivalService.getAllFestivals().subscribe({
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
    this.router.navigate(['festivals/', festival.id]);
  }

  get filteredFestivals(): Festival[] {
    if (this.selectedChip === 'Upcoming') {
      return this.festivals.filter((festival) => festival.status === 'PUBLIC');
    } else if (this.selectedChip === 'Past') {
      return this.festivals.filter(
        (festival) => festival.status === 'COMPLETED',
      );
    } else {
      return this.festivals;
    }
  }
}
