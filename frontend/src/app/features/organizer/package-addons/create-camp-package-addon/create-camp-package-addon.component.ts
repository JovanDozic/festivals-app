import { Component, inject, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import {
  FormArray,
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import {
  AddCampConfigRequest,
  CreateItemPriceRequest,
  CreateItemRequest,
  EquipmentDTO,
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
import { ItemService } from '../../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSelectModule } from '@angular/material/select';
import { MatTimepickerModule } from '@angular/material/timepicker';
import { ImageService } from '../../../../services/image/image.service';

@Component({
  selector: 'app-create-camp-package-addon',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatTimepickerModule,
    MatGridListModule,
    MatIconModule,
    MatTabsModule,
    MatStepperModule,
    MatSlideToggleModule,
    MatDialogModule,
    MatSelectModule,
  ],
  templateUrl: './create-camp-package-addon.component.html',
  styleUrls: [
    './create-camp-package-addon.component.scss',
    '../../../../app.component.scss',
  ],
})
export class CreateCampPackageAddonComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<CreateCampPackageAddonComponent>);
  private data: { festivalId: number; category: string } =
    inject(MAT_DIALOG_DATA);
  private itemService = inject(ItemService);
  private imageService = inject(ImageService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  category = '';

  infoFormGroup: FormGroup;
  configurationFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;

  itemId: number | null = null;
  isFixedPrice = true;

  selectedFile: File | null = null;
  imagePreviewUrl: string | ArrayBuffer | null = null;

  constructor() {
    this.category = this.data?.category;

    this.infoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      availableNumberCtrl: ['', [Validators.required, Validators.min(1)]],
    });

    this.configurationFormGroup = this.fb.group({
      campNameCtrl: ['', Validators.required],
      equipmentFormArray: this.fb.array([this.createEquipmentFormGroup()]),
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  get equipmentFormArray(): FormArray {
    return this.configurationFormGroup.get('equipmentFormArray') as FormArray;
  }

  private createEquipmentFormGroup(): FormGroup {
    return this.fb.group({
      equipmentNameCtrl: ['', Validators.required],
    });
  }

  addEquipment() {
    const lastGroup = this.equipmentFormArray.at(
      this.equipmentFormArray.length - 1,
    ) as FormGroup;

    if (lastGroup.valid) {
      this.equipmentFormArray.push(this.createEquipmentFormGroup());
    } else {
      this.snackbarService.show('Please fill out the last equipment');
    }
  }

  removeEquipment(index: number) {
    if (this.equipmentFormArray.length > 1) {
      this.equipmentFormArray.removeAt(index);
    } else {
      this.snackbarService.show('At least one equipment is required');
    }
  }

  done() {
    if (this.itemId) {
      this.createFixedPrice();
    }
  }

  createPackageAddon() {
    if (this.infoFormGroup.valid && this.data.festivalId) {
      const request: CreateItemRequest = {
        name: this.infoFormGroup.get('nameCtrl')?.value,
        description: this.infoFormGroup.get('descriptionCtrl')?.value,
        availableNumber: this.infoFormGroup.get('availableNumberCtrl')?.value,
        type: 'PACKAGE_ADDON',
      };

      this.itemService
        .createPackageAddon(this.data.festivalId, request, this.category)
        .subscribe({
          next: (response) => {
            this.snackbarService.show('Package Addon created');
            this.itemId = response;
            this.stepper?.next();
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error creating Package Addon');
          },
        });
    }
  }

  addCampConfig() {
    if (this.configurationFormGroup.valid && this.itemId && this.selectedFile) {
      const equipmentList: EquipmentDTO[] = [];
      this.equipmentFormArray.controls.forEach((control) => {
        const equipmentName = control.get('equipmentNameCtrl')?.value;
        if (equipmentName) {
          equipmentList.push({ name: equipmentName });
        }
      });

      const request: AddCampConfigRequest = {
        itemId: this.itemId,
        campName: this.configurationFormGroup.get('campNameCtrl')?.value,
        imageURL: '',
        equipmentList: equipmentList,
      };

      this.imageService.uploadImageAndGetURL(this.selectedFile).subscribe({
        next: (response) => {
          request.imageURL = response.imageURL;

          this.itemService
            .addCampConfig(this.data.festivalId, request)
            .subscribe({
              next: () => {
                this.snackbarService.show('Camp Config created');
                this.stepper?.next();
              },
              error: (error) => {
                console.log(error);
                this.snackbarService.show('Error creating Camp Config');
              },
            });
        },
      });
    }
  }

  createFixedPrice() {
    if (this.fixedPriceFormGroup.valid && this.itemId && this.data.festivalId) {
      const request: CreateItemPriceRequest = {
        itemId: this.itemId,
        price: this.fixedPriceFormGroup.get('fixedPriceCtrl')?.value,
        isFixed: true,
      };

      this.itemService
        .createItemPrice(this.data.festivalId, request)
        .subscribe({
          next: () => {
            this.snackbarService.show('Fixed Price created');
            this.dialogRef.close(true);
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error creating Fixed Price');
            this.dialogRef.close(false);
          },
        });
    }
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
}
