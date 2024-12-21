import { Component, inject, OnInit } from '@angular/core';
import {
  Festival,
  TransportAddonDTO,
} from '../../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../../services/festival/festival.service';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { ItemService } from '../../../../services/festival/item.service';
import { CreateTransportPackageAddonComponent } from '../create-transport-package-addon/create-transport-package-addon.component';

@Component({
  selector: 'app-transport-package-addons',
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
  templateUrl: './transport-package-addons.component.html',
  styleUrls: [
    './transport-package-addons.component.scss',
    '../../../../app.component.scss',
  ],
})
export class TransportPackageAddonsComponent implements OnInit {
  isLoading = true;
  festival: Festival | null = null;
  transportCount = 0;

  transportAddons: TransportAddonDTO[] = [];

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private itemService = inject(ItemService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadAddons();
  }

  loadAddons() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getTransportAddons(Number(id)).subscribe({
        next: (response) => {
          this.transportAddons = response;
          this.transportCount = this.transportAddons.length;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting Travel Addons');
          this.transportAddons = [];
        },
      });
    }
  }

  goBack() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/package-addons`,
    ]);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
          this.isLoading = true;
        },
      });
    }
  }

  onAddAddonClick() {
    const dialogRef = this.dialog.open(CreateTransportPackageAddonComponent, {
      data: { festivalId: this.festival?.id, category: 'TRANSPORT' },
      width: '800px',
      height: '700px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadAddons();
      }
    });
  }
}
