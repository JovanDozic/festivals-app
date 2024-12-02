import { Component, inject, OnInit } from '@angular/core';
import { Festival } from '../../../../models/festival/festival.model';
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
import { CreatePackageAddonChooserComponent } from '../create-package-addon-chooser/create-package-addon-chooser.component';
import { CreateTransportPackageAddonComponent } from '../create-transport-package-addon/create-transport-package-addon.component';
import { CreateCampPackageAddonComponent } from '../create-camp-package-addon/create-camp-package-addon.component';

@Component({
  selector: 'app-package-addons',
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
  templateUrl: './package-addons.component.html',
  styleUrls: [
    './package-addons.component.scss',
    '../../../../app.component.scss',
  ],
})
export class PackageAddonsComponent implements OnInit {
  isLoading: boolean = true;
  festival: Festival | null = null;
  generalCount: number = 0;
  transportCount: number = 0;
  campCount: number = 0;

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private itemService = inject(ItemService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadCounts();
  }

  loadCounts() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getPackageAddonsCount(Number(id), 'general').subscribe({
        next: (count) => {
          this.generalCount = count;
        },
        error: (error) => {
          console.log('Error fetching general package addon count: ', error);
          this.snackbarService.show(
            'Error getting general package addon count'
          );
          this.generalCount = 0;
        },
      });

      this.itemService
        .getPackageAddonsCount(Number(id), 'transport')
        .subscribe({
          next: (count) => {
            this.transportCount = count;
          },
          error: (error) => {
            console.log(
              'Error fetching transport package addon count: ',
              error
            );
            this.snackbarService.show(
              'Error getting transport package addon count'
            );
            this.transportCount = 0;
          },
        });

      this.itemService.getPackageAddonsCount(Number(id), 'camp').subscribe({
        next: (count) => {
          this.campCount = count;
        },
        error: (error) => {
          console.log('Error fetching camp package addon count: ', error);
          this.snackbarService.show('Error getting camp package addon count');
          this.campCount = 0;
        },
      });
    }
  }

  goBack() {
    this.router.navigate([`organizer/my-festivals/${this.festival?.id}`]);
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

  onViewGeneralAddonsClick() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/package-addons/general`,
    ]);
  }

  onViewTransportAddonsClick() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/package-addons/transport`,
    ]);
  }

  onViewCampAddonsClick() {
    this.router.navigate([
      `organizer/my-festivals/${this.festival?.id}/package-addons/camp`,
    ]);
  }

  onAddPackageAddonClick() {
    const dialogRef = this.dialog.open(CreatePackageAddonChooserComponent, {
      data: { festivalId: this.festival?.id },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.launchCreatePackageAddonDialog(result);
      }
    });
  }

  launchCreatePackageAddonDialog(result: any) {
    if (!result) {
      return;
    }

    if (result === 'GENERAL') {
      this.launchCreateGeneralPackageAddonDialog(result);
    }

    if (result === 'TRANSPORT') {
      this.launchCreateTransportPackageAddonDialog(result);
    }

    if (result === 'CAMP') {
      this.launchCreateCampPackageAddonDialog(result);
    }
  }

  launchCreateCampPackageAddonDialog(result: any) {
    const dialogRef = this.dialog.open(CreateCampPackageAddonComponent, {
      data: { festivalId: this.festival?.id, category: result },
      width: '800px',
      height: '700px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadCounts();
      }
    });
  }

  launchCreateTransportPackageAddonDialog(result: any) {
    const dialogRef = this.dialog.open(CreateTransportPackageAddonComponent, {
      data: { festivalId: this.festival?.id, category: result },
      width: '800px',
      height: '700px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadCounts();
      }
    });
  }

  launchCreateGeneralPackageAddonDialog(result: any) {
    const dialogRef = this.dialog.open(CreateGeneralPackageAddonComponent, {
      data: { festivalId: this.festival?.id, category: result },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadCounts();
      }
    });
  }
}
