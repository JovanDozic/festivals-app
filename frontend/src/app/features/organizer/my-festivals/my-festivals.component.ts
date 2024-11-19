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

  ngOnInit(): void {
    this.loadFestivals();
  }

  loadFestivals(): void {
    this.festivalService.getMyFestivals().subscribe({
      next: (response) => {
        setTimeout(() => {
          console.log('Festivals:', response);
          this.festivals = response;
          this.isLoading = false;
        }, 2000);
      },
      error: (error) => {
        console.error('Error fetching festivals:', error);
        this.snackbarService.show('Error fetching festivals');
        // this.isLoading = false;
      },
    });
  }
}
