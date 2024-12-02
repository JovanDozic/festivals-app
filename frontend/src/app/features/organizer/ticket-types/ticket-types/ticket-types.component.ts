import { Component, inject, OnInit } from '@angular/core';
import {
  Festival,
  ItemCurrentPrice,
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
import { MatTableModule } from '@angular/material/table';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../../shared/confirmation-dialog/confirmation-dialog.component';
import { ItemService } from '../../../../services/festival/item.service';
import { CreateTicketTypeComponent } from '../create-ticket-type/create-ticket-type.component';
import { ViewEditTicketTypeComponent } from '../view-edit-ticket-type/view-edit-ticket-type.component';

@Component({
  selector: 'app-ticket-types',
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
  ],
  templateUrl: './ticket-types.component.html',
  styleUrls: [
    './ticket-types.component.scss',
    '../../../../app.component.scss',
  ],
})
export class TicketTypesComponent implements OnInit {
  festival: Festival | null = null;
  ticketTypesCount = 0;
  ticketTypes: ItemCurrentPrice[] = [];
  displayedColumns = ['id', 'name', 'price', 'isFixed', 'dateTo', 'actions'];

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private itemService = inject(ItemService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadTicketTypes();
  }

  goBack() {
    this.router.navigate([`organizer/my-festivals/${this.festival?.id}`]);
  }

  loadTicketTypes() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getTicketTypes(Number(id)).subscribe({
        next: (ticketTypes) => {
          this.ticketTypes = ticketTypes;
          this.ticketTypesCount = this.ticketTypes.length;
        },
        error: (error) => {
          console.log('Error fetching ticket types: ', error);
          this.snackbarService.show('Error getting ticket types');
          this.ticketTypes = [];
          this.ticketTypesCount = 0;
        },
      });
    }
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
        },
        error: (error) => {
          console.log('Error fetching festival information: ', error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
        },
      });
    }
  }

  onCreateTicketTypeClick() {
    const dialogRef = this.dialog.open(CreateTicketTypeComponent, {
      data: {
        festivalId: this.festival?.id,
      },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadTicketTypes();
      }
    });
  }

  onViewClick(itemId: number) {
    const dialogRef = this.dialog.open(ViewEditTicketTypeComponent, {
      data: {
        festivalId: this.festival?.id,
        itemId: itemId,
      },
      width: '800px',
      height: 'auto',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadTicketTypes();
      }
    });
  }

  onDeleteTicketTypeClick(itemId: number, name: string) {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Fire Employee',
        message: `Are you sure you want to delete ${name}? You won't be able to delete this ticket type if it has been used in any transactions.`,
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.itemService
          .deleteItem(Number(this.festival?.id), itemId)
          .subscribe({
            next: () => {
              this.snackbarService.show('Ticket type deleted');
              this.loadTicketTypes();
            },
            error: (error) => {
              console.log('Error deleting ticket type: ', error);
              this.snackbarService.show('Error deleting ticket type');
            },
          });
      }
    });
  }
}
