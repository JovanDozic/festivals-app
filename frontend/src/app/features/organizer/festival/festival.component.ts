import { Component, inject, OnInit } from '@angular/core';
import { Festival } from '../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../services/festival/festival.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';

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
  ],
  templateUrl: './festival.component.html',
  styleUrl: './festival.component.scss',
})
export class FestivalComponent implements OnInit {
  festival: Festival | null = null;
  isLoading: boolean = true;

  private route = inject(ActivatedRoute);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);

  ngOnInit() {
    this.getFestival();
  }

  getFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          console.log('Festival: ', festival);
          this.festival = festival;
          console.log('this.festival.name', this.festival.name);
          this.isLoading = false;
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
}
