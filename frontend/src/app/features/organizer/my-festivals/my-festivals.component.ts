import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { Festival } from '../../../models/festival/festival.model';
import { FestivalService } from '../../../services/festival/festival.service';

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
  ],
})
export class MyFestivalsComponent implements OnInit {
  festivals: Festival[] = [];

  private festivalService = inject(FestivalService);

  ngOnInit(): void {
    this.loadFestivals();
  }

  loadFestivals(): void {
    this.festivalService.getMyFestivals().subscribe({
      next: (response) => {
        this.festivals = response;
        console.log('Response: ', response);
        console.log('Festivals: ', response);
        console.log('My Festivals: ', this.festivals);
      },
      error: (error) => {
        console.error('Error fetching festivals:', error);
      },
    });
  }
}
