import { Component, inject, OnInit } from '@angular/core';
import { Festival } from '../../../models/festival/festival.model';
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
import { animate, style, transition, trigger } from '@angular/animations';
import { MatMenuModule } from '@angular/material/menu';
import { ItemService } from '../../../services/festival/item.service';
import { StoreChooserComponent } from '../store/store-chooser/store-chooser.component';

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
    MatChipsModule,
    MatMenuModule,
  ],
  templateUrl: './festival.component.html',
  styleUrls: ['./festival.component.scss', '../../../app.component.scss'],
  animations: [
    trigger('fadeAnimation', [
      transition(':increment', [
        style({ opacity: 0 }),
        animate('500ms ease-in', style({ opacity: 1 })),
      ]),
      transition(':decrement', [
        style({ opacity: 0 }),
        animate('500ms ease-in', style({ opacity: 1 })),
      ]),
    ]),
  ],
})
export class FestivalComponent implements OnInit {
  festival: Festival | null = null;
  isLoading = true;
  currentImageIndex = 0;
  previousImageIndex = 0;
  ticketTypesCount = 0;
  packageAddonsCount = 0;

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private itemService = inject(ItemService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadTicketTypesCount();
    this.loadPackageAddonsCount();
  }

  goBack() {
    this.router.navigate(['festivals']);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
          this.isLoading = false;
          this.currentImageIndex = 0;
          this.previousImageIndex = 0;
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

  loadTicketTypesCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getTicketTypesCount(Number(id)).subscribe({
        next: (count) => {
          this.ticketTypesCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting ticket types count');
          this.ticketTypesCount = 0;
        },
      });
    }
  }

  loadPackageAddonsCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getAllPackageAddonsCount(Number(id)).subscribe({
        next: (count) => {
          this.packageAddonsCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting package addons count');
          this.packageAddonsCount = 0;
        },
      });
    }
  }

  nextImage() {
    if (this.festival && this.festival.images) {
      this.previousImageIndex = this.currentImageIndex;
      this.currentImageIndex =
        (this.currentImageIndex + 1) % this.festival.images.length;
    }
  }

  previousImage() {
    if (this.festival && this.festival.images) {
      this.previousImageIndex = this.currentImageIndex;
      this.currentImageIndex =
        (this.currentImageIndex - 1 + this.festival.images.length) %
        this.festival.images.length;
    }
  }

  onOpenStoreClick() {
    if (this.ticketTypesCount === 0 && this.packageAddonsCount === 0) {
      this.snackbarService.show('No items available for sale yet');
      return;
    }

    if (this.ticketTypesCount > 0 && this.packageAddonsCount === 0) {
      this.router.navigate([
        'festivals/' + this.festival?.id + '/store/ticket',
      ]);
      return;
    }

    const dialogRef = this.dialog.open(StoreChooserComponent, {
      data: { festivalId: this.festival?.id },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.router.navigate([
          'festivals/' + this.festival?.id + '/store/' + result,
        ]);
      }
    });
  }
}
