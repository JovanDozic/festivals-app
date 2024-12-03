import { Component, inject, OnInit } from '@angular/core';
import {
  Employee,
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
import { MatStepperModule } from '@angular/material/stepper';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-store-ticket',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
    MatTableModule,
    MatInputModule,
    MatFormFieldModule,
    MatStepperModule,
  ],
  templateUrl: './store-ticket.component.html',
  styleUrls: [
    './store-ticket.component.scss',
    '../../../../app.component.scss',
  ],
})
export class StoreTicketComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);
  private fb = inject(FormBuilder);

  festival: Festival | null = null;
  isLoading = false;

  infoFormGroup: FormGroup;
  tickets: ItemCurrentPrice[] = [
    {
      itemId: 1,
      name: 'General Admission',
      description: 'General Admission Ticket',
      price: 50.0,
      remainingNumber: 100,
      availableNumber: 100,
      isFixed: true,
      dateFrom: '2022-06-01',
      dateTo: '2022-06-30',
      priceListItemId: 1,
    },
    {
      itemId: 2,
      name: 'VIP Admission',
      description: 'VIP Admission Ticket',
      price: 100.0,
      remainingNumber: 50,
      availableNumber: 50,
      isFixed: true,
      dateFrom: '2022-06-01',
      dateTo: '2022-06-30',
      priceListItemId: 2,
    },
    {
      itemId: 2,
      name: 'VIP Admission',
      description: 'VIP Admission Ticket',
      price: 100.0,
      remainingNumber: 50,
      availableNumber: 50,
      isFixed: true,
      dateFrom: '2022-06-01',
      dateTo: '2022-06-30',
      priceListItemId: 2,
    },
    {
      itemId: 2,
      name: 'VIP Admission',
      description: 'VIP Admission Ticket',
      price: 100.0,
      remainingNumber: 50,
      availableNumber: 50,
      isFixed: true,
      dateFrom: '2022-06-01',
      dateTo: '2022-06-30',
      priceListItemId: 2,
    },
    {
      itemId: 2,
      name: 'VIP Admission',
      description: 'VIP Admission Ticket',
      price: 100.0,
      remainingNumber: 50,
      availableNumber: 50,
      isFixed: true,
      dateFrom: '2022-06-01',
      dateTo: '2022-06-30',
      priceListItemId: 2,
    },
  ];
  selectedTicket: ItemCurrentPrice | null = null;

  constructor() {
    this.infoFormGroup = this.fb.group({
      nameCtrl: [''],
    });
  }

  ngOnInit() {
    this.loadFestival();
  }

  goBack() {
    this.router.navigate([`festivals/${this.festival?.id}`]);
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
          console.log('Error fetching festival information: ', error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
          this.isLoading = true;
        },
      });
    }
  }

  purchase() {
    this.snackbarService.show('Opening Payment Dialog...');
  }

  selectTicket(ticket: ItemCurrentPrice) {
    this.selectedTicket = ticket;
    console.log('AAALO', ticket);
    console.log('AAALO', this.selectedTicket);
  }
}
