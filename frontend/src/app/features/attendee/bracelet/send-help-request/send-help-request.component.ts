import { Component, inject } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import {
  ActivateBraceletHelpRequest,
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
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { OrderService } from '../../../../services/festival/order.service';
import { ImageService } from '../../../../services/image/image.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

@Component({
  selector: 'app-send-help-request',
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
  templateUrl: './send-help-request.component.html',
  styleUrls: [
    './send-help-request.component.scss',
    '../../../../app.component.scss',
  ],
})
export class SendHelpRequestComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<SendHelpRequestComponent>);
  private data: {
    order: OrderDTO;
  } = inject(MAT_DIALOG_DATA);
  private orderService = inject(OrderService);
  private imageService = inject(ImageService);

  infoFormGroup: FormGroup;
  selectedFile: File | null = null;
  imagePreviewUrl: string | ArrayBuffer | null = null;
  isUploading = false;

  constructor() {
    this.infoFormGroup = this.fb.group({
      pinUserCtrl: [
        '',
        [
          Validators.minLength(4),
          Validators.maxLength(20),
          Validators.pattern('^[0-9]+$'),
        ],
      ],
      barcodeNumberSystemCtrl: this.data.order.bracelet?.barcodeNumber,
      barcodeNumberUserCtrl: ['', Validators.maxLength(20)],
      issueDescriptionCtrl: [
        '',
        [
          Validators.required,
          Validators.minLength(5),
          Validators.maxLength(200),
        ],
      ],
    });

    this.infoFormGroup.get('barcodeNumberSystemCtrl')?.disable();
  }

  onFileSelected(event: Event) {
    const fileInput = event.target as HTMLInputElement;

    if (fileInput.files && fileInput.files.length > 0) {
      this.selectedFile = fileInput.files[0];

      const reader = new FileReader();

      reader.onload = () => {
        this.imagePreviewUrl = reader.result;
      };

      reader.readAsDataURL(this.selectedFile);
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  sendRequest() {
    if (
      this.infoFormGroup.valid &&
      this.data.order &&
      this.selectedFile &&
      this.data.order.bracelet
    ) {
      this.isUploading = true;
      const request: ActivateBraceletHelpRequest = {
        braceletId: this.data.order.bracelet.braceletId,
        pinUser: this.infoFormGroup.get('pinUserCtrl')?.value,
        barcodeNumberUser: this.infoFormGroup.get('barcodeNumberUserCtrl')
          ?.value,
        issueDescription: this.infoFormGroup.get('issueDescriptionCtrl')?.value,
        imageURL: '',
      };

      this.imageService.uploadImageAndGetURL(this.selectedFile).subscribe({
        next: (response) => {
          request.imageURL = response.imageURL;
          this.orderService.sendHelpRequest(request).subscribe({
            next: () => {
              this.snackbarService.show('Help Request sent successfully');
              this.isUploading = false;
              this.dialogRef.close(true);
            },
            error: (error) => {
              console.log(error);
              this.snackbarService.show('Error sending Help Request');
              this.isUploading = false;
            },
          });
        },
      });
    }
  }
}
