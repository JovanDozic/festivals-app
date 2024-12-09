import { Component, inject, OnInit } from '@angular/core';
import {
  Festival,
  GeneralAddonDTO,
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
import { CreateGeneralPackageAddonComponent } from '../create-general-package-addon/create-general-package-addon.component';

@Component({
  selector: 'app-general-package-addons',
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
  templateUrl: './general-package-addons.component.html',
  styleUrls: [
    './general-package-addons.component.scss',
    '../../../../app.component.scss',
  ],
})
export class GeneralPackageAddonsComponent implements OnInit {
  isLoading = true;
  festival: Festival | null = null;
  generalCount = 0;

  generalAddons: GeneralAddonDTO[] = [];

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
      this.itemService.getGeneralAddons(Number(id)).subscribe({
        next: (response) => {
          this.generalAddons = response;
          this.generalCount = this.generalAddons.length;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting General Addons');
          this.generalAddons = [];
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
    const dialogRef = this.dialog.open(CreateGeneralPackageAddonComponent, {
      data: { festivalId: this.festival?.id, category: 'GENERAL' },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadAddons();
      }
    });
  }
}
