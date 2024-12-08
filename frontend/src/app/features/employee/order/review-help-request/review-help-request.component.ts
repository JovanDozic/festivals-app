import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  NumberValueAccessor,
  ReactiveFormsModule,
} from '@angular/forms';
import {
  ActivationHelpRequestDTO,
  OrderDTO,
} from '../../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { MatStepperModule } from '@angular/material/stepper';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { OrderService } from '../../../../services/festival/order.service';
import { ImageService } from '../../../../services/image/image.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ItemService } from '../../../../services/festival/item.service';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-review-help-request',
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
    MatTabsModule,
    MatStepperModule,
    MatSlideToggleModule,
    MatDialogModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './review-help-request.component.html',
  styleUrls: [
    './review-help-request.component.scss',
    '../../../../app.component.scss',
  ],
})
export class ReviewHelpRequestComponent implements OnInit {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ReviewHelpRequestComponent>);
  private data: {
    festivalId: number;
    orderId: number;
    festivalTicketId: number;
    attendeeUsername: string;
    braceletId: number;
  } = inject(MAT_DIALOG_DATA);
  private orderService = inject(OrderService);
  private itemService = inject(ItemService);
  private imageService = inject(ImageService);

  order: OrderDTO | null = null;
  helpRequest: ActivationHelpRequestDTO | null = null;

  constructor() {}

  ngOnInit(): void {
    this.loadOrder();
    this.loadHelpRequest();
  }

  loadHelpRequest() {
    if (this.data.braceletId) {
      this.orderService.getHelpRequest(this.data.braceletId).subscribe({
        next: (helpRequest) => {
          console.log(helpRequest);
          this.helpRequest = helpRequest;
        },
        error: (error) => {
          console.log(error);
          if (error.status === 404) {
            this.snackbarService.show('Help Request not found');
          } else {
            this.snackbarService.show('Error loading Help Request');
          }
        },
      });
    }
  }

  loadOrder() {
    if (this.data.orderId) {
      this.itemService.getOrder(this.data.orderId).subscribe({
        next: (order) => {
          console.log(order);
          this.order = order;
        },
        error: (error) => {
          console.log(error);
          if (error.status === 404) {
            this.snackbarService.show('Order not found');
          } else {
            this.snackbarService.show('Error loading order');
          }
        },
      });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  saveChanges() {}
}
