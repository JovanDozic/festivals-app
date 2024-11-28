import { Component, inject, OnInit, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormArray,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { FestivalService } from '../../../services/festival/festival.service';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
  Item,
  VariablePrice,
} from '../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../services/user/user.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { ItemService } from '../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-view-edit-ticket-type',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatGridListModule,
    MatIconModule,
    MatDialogTitle,
    MatDialogContent,
    MatTabsModule,
    MatStepperModule,
    MatSlideToggleModule,
  ],
  templateUrl: './view-edit-ticket-type.component.html',
  styleUrls: [
    './view-edit-ticket-type.component.scss',
    '../../../app.component.scss',
  ],
})
export class ViewEditTicketTypeComponent implements OnInit {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ViewEditTicketTypeComponent>);
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);
  private itemService = inject(ItemService);

  isEditing: boolean = false;

  infoFormGroup: FormGroup;
  priceFormGroup: FormGroup;
  variablePriceFormGroup: FormGroup;

  ticketType: Item | null = null;

  ngOnInit(): void {
    this.loadTicketType();
  }

  constructor() {
    // todo: fill in information for selected ticket type
    // todo: fetch ticket type information from server
    this.infoFormGroup = this.fb.group({
      name: ['', Validators.required],
      description: ['', Validators.required],
      availableNumber: [0, Validators.required],
      type: [''],
    });

    this.priceFormGroup = this.fb.group({
      price: [0, Validators.required],
      isFixed: [true, Validators.required],
    });

    this.variablePriceFormGroup = this.fb.group({
      variablePrices: this.fb.array([]),
    });
  }

  loadTicketType() {
    const festivalId = this.data.festivalId;
    const ticketTypeId = 29; // todo: get ticket type id from dialog data
    if (festivalId && ticketTypeId) {
      this.itemService.getTicketType(festivalId, ticketTypeId).subscribe({
        next: (ticketType) => {
          this.ticketType = ticketType;
          console.log('Ticket type: ', ticketType);
        },
        error: (error) => {
          console.log('Error fetching ticket type: ', error);
          this.snackbarService.show('Error getting ticket type');
        },
      });
    }
  }
}
