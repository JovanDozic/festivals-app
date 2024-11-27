import { Component, inject, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { FestivalService } from '../../../services/festival/festival.service';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
} from '../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import {
  MATERIAL_SANITY_CHECKS,
  provideNativeDateAdapter,
} from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../services/user/user.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { ItemService } from '../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

@Component({
  selector: 'app-create-ticket-type',
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
    MatDialogActions,
    MatTabsModule,
    MatStepperModule,
    MatSlideToggleModule,
  ],
  templateUrl: './create-ticket-type.component.html',
  styleUrls: [
    './create-ticket-type.component.scss',
    '../../../app.component.scss',
  ],
})
export class CreateTicketTypeComponent {
  private fb = inject(FormBuilder);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<CreateTicketTypeComponent>);
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);
  private userService = inject(UserService);
  private itemService = inject(ItemService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  infoFormGroup: FormGroup;
  ticketTypeId: number | null = null;
  isFixedPrice: boolean = true;

  fixedPriceFormGroup: FormGroup;

  constructor() {
    this.infoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      availableNumberCtrl: ['', [Validators.required, Validators.min(1)]],
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });
  }

  toggleIsFixed() {
    this.isFixedPrice = !this.isFixedPrice;
  }

  createTicketType() {
    if (this.infoFormGroup.valid && this.data.festivalId) {
      const request: CreateItemRequest = {
        name: this.infoFormGroup.get('nameCtrl')?.value,
        description: this.infoFormGroup.get('descriptionCtrl')?.value,
        availableNumber: this.infoFormGroup.get('availableNumberCtrl')?.value,
        type: 'TICKET_TYPE',
      };

      this.itemService.createItem(this.data.festivalId, request).subscribe({
        next: (response) => {
          console.log('Ticket type created: ', response);
          this.snackbarService.show('Ticket Type created');
          this.ticketTypeId = response;
          this.stepper?.next();
        },
        error: (error) => {
          console.log('Error creating ticket type: ', error);
          this.snackbarService.show('Error creating Ticket Type');
        },
      });
    }
  }

  done() {
    if (this.ticketTypeId) {
      if (this.isFixedPrice) {
        this.createFixedPrice();
      } else {
        this.createNotFixedPrices();
      }
    }
  }

  createFixedPrice() {
    if (
      this.fixedPriceFormGroup.valid &&
      this.ticketTypeId &&
      this.data.festivalId
    ) {
      const request: CreateItemPriceRequest = {
        itemId: this.ticketTypeId,
        price: this.fixedPriceFormGroup.get('fixedPriceCtrl')?.value,
        isFixed: true,
      };

      this.itemService
        .createItemPrice(this.data.festivalId, request)
        .subscribe({
          next: (response) => {
            console.log('Fixed price created: ', response);
            this.snackbarService.show('Fixed Price created');
            this.dialogRef.close(true);
          },
          error: (error) => {
            console.log('Error creating fixed price: ', error);
            this.snackbarService.show('Error creating Fixed Price');
            // todo: uncomment
            // this.dialogRef.close(false);
          },
        });
    }
  }

  createNotFixedPrices() {}
}
