import { Component, inject, OnInit } from '@angular/core';
import {
  CampAddonDTO,
  Festival,
} from '../../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../../services/festival/festival.service';
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
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
import { CreateCampPackageAddonComponent } from '../create-camp-package-addon/create-camp-package-addon.component';
@Component({
  selector: 'app-camp-package-addons',
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
  templateUrl: './camp-package-addons.component.html',
  styleUrls: [
    './camp-package-addons.component.scss',
    '../../../../app.component.scss',
  ],
})
export class CampPackageAddonsComponent implements OnInit {
  isLoading = true;
  festival: Festival | null = null;
  campCount = 0;

  campAddons: CampAddonDTO[] = [];

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
      this.itemService.getCampAddons(Number(id)).subscribe({
        next: (response) => {
          console.log(`General Addons`, response);
          this.campAddons = response;
          this.campCount = this.campAddons.length;
        },
        error: (error) => {
          console.log('Error fetching general addons: ', error);
          this.snackbarService.show('Error getting General Addons');
          this.campAddons = [];
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
          console.log(
            `Festival ID: <${this.festival?.id}> - ${this.festival?.name}`
          );
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

  onAddAddonClick() {
    const dialogRef = this.dialog.open(CreateCampPackageAddonComponent, {
      data: { festivalId: this.festival?.id, category: 'CAMP' },
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
