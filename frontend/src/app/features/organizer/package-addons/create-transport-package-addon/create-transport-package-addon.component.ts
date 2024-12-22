import { Component, inject, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import {
  AddTransportConfigRequest,
  CreateItemPriceRequest,
  CreateItemRequest,
  TransportType,
} from '../../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import {
  DateAdapter,
  MAT_DATE_FORMATS,
  MAT_DATE_LOCALE,
} from '@angular/material/core';
import { CustomDateAdapter } from '../../../../shared/date-formats/date-adapter';
import { CUSTOM_DATE_FORMATS } from '../../../../shared/date-formats/date-formats';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { ItemService } from '../../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSelectModule } from '@angular/material/select';
import { MatTimepickerModule } from '@angular/material/timepicker';
import { CityRequest } from '../../../../models/common/address.model';
import { CountryPickerComponent } from '../../../../shared/country-picker/country-picker.component';

@Component({
  selector: 'app-create-transport-package-addon',
  imports: [
    CountryPickerComponent,
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
  templateUrl: './create-transport-package-addon.component.html',
  styleUrls: [
    './create-transport-package-addon.component.scss',
    '../../../../app.component.scss',
  ],
  providers: [
    { provide: DateAdapter, useClass: CustomDateAdapter },
    { provide: MAT_DATE_FORMATS, useValue: CUSTOM_DATE_FORMATS },
    { provide: MAT_DATE_LOCALE, useValue: 'en-GB' },
  ],
})
export class CreateTransportPackageAddonComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(
    MatDialogRef<CreateTransportPackageAddonComponent>,
  );
  private data: { festivalId: number; category: string } =
    inject(MAT_DIALOG_DATA);
  private itemService = inject(ItemService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  category = '';

  infoFormGroup: FormGroup;
  configurationFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;

  itemId: number | null = null;
  isFixedPrice = true;

  transportTypes: TransportType[] = [
    { value: 'BUS', viewValue: 'Bus' },
    { value: 'PLANE', viewValue: 'Plane' },
    { value: 'TRAIN', viewValue: 'Train' },
  ];

  selectedTransportType: TransportType | null = null;

  constructor() {
    this.category = this.data?.category;

    this.infoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      availableNumberCtrl: ['', [Validators.required, Validators.min(1)]],
    });

    this.configurationFormGroup = this.fb.group({
      transportTypeCtrl: ['', Validators.required],
      departureCityNameCtrl: ['', Validators.required],
      departureCityPostalCodeCtrl: ['', Validators.required],
      departureCountryISO3Ctrl: [
        '',
        [Validators.required, Validators.maxLength(3)],
      ],
      arrivalCityNameCtrl: ['', Validators.required],
      arrivalCityPostalCodeCtrl: ['', Validators.required],
      arrivalCountryISO3Ctrl: [
        '',
        [Validators.required, Validators.maxLength(3)],
      ],

      departureDateCtrl: ['', Validators.required],
      departureTimeCtrl: ['', Validators.required],

      arrivalDateCtrl: ['', Validators.required],
      arrivalTimeCtrl: ['', Validators.required],

      returnDepartureDateCtrl: ['', Validators.required],
      returnDepartureTimeCtrl: ['', Validators.required],

      returnArrivalDateCtrl: ['', Validators.required],
      returnArrivalTimeCtrl: ['', Validators.required],
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
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

  addTransportConfig() {
    if (this.configurationFormGroup.valid && this.itemId) {
      const departureDate: Date =
        this.configurationFormGroup.get('departureDateCtrl')?.value;
      const departureTime: Date =
        this.configurationFormGroup.get('departureTimeCtrl')?.value;
      const arrivalDate: Date =
        this.configurationFormGroup.get('arrivalDateCtrl')?.value;
      const arrivalTime: Date =
        this.configurationFormGroup.get('arrivalTimeCtrl')?.value;
      const returnDepartureDate: Date = this.configurationFormGroup.get(
        'returnDepartureDateCtrl',
      )?.value;
      const returnDepartureTime: Date = this.configurationFormGroup.get(
        'returnDepartureTimeCtrl',
      )?.value;
      const returnArrivalDate: Date = this.configurationFormGroup.get(
        'returnArrivalDateCtrl',
      )?.value;
      const returnArrivalTime: Date = this.configurationFormGroup.get(
        'returnArrivalTimeCtrl',
      )?.value;

      const departureDateTime = this.combineDateAndTime(
        departureDate,
        departureTime,
      );
      const arrivalDateTime = this.combineDateAndTime(arrivalDate, arrivalTime);
      const returnDepartureDateTime = this.combineDateAndTime(
        returnDepartureDate,
        returnDepartureTime,
      );
      const returnArrivalDateTime = this.combineDateAndTime(
        returnArrivalDate,
        returnArrivalTime,
      );

      const departureDateTimeFormatted = formatDateTime(departureDateTime);
      const arrivalDateTimeFormatted = formatDateTime(arrivalDateTime);
      const returnDepartureDateTimeFormatted = formatDateTime(
        returnDepartureDateTime,
      );
      const returnArrivalDateTimeFormatted = formatDateTime(
        returnArrivalDateTime,
      );

      const departureCity: CityRequest = {
        name: this.configurationFormGroup.get('departureCityNameCtrl')?.value,
        postalCode: this.configurationFormGroup.get(
          'departureCityPostalCodeCtrl',
        )?.value,
        countryISO3: this.configurationFormGroup.get('departureCountryISO3Ctrl')
          ?.value,
      };

      const arrivalCity: CityRequest = {
        name: this.configurationFormGroup.get('arrivalCityNameCtrl')?.value,
        postalCode: this.configurationFormGroup.get('arrivalCityPostalCodeCtrl')
          ?.value,
        countryISO3: this.configurationFormGroup.get('arrivalCountryISO3Ctrl')
          ?.value,
      };

      const request: AddTransportConfigRequest = {
        itemId: this.itemId,
        transportType:
          this.configurationFormGroup.get('transportTypeCtrl')?.value,
        departureCity: departureCity,
        arrivalCity: arrivalCity,
        departureTime: departureDateTimeFormatted,
        arrivalTime: arrivalDateTimeFormatted,
        returnDepartureTime: returnDepartureDateTimeFormatted,
        returnArrivalTime: returnArrivalDateTimeFormatted,
      };

      this.itemService
        .addTransportConfig(this.data.festivalId, request)
        .subscribe({
          next: () => {
            this.snackbarService.show('Transport configuration added');
            this.stepper?.next();
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error adding transport configuration');
          },
        });
    }
  }

  private combineDateAndTime(date: Date, time: Date): Date {
    const hours = time.getHours();
    const minutes = time.getMinutes();
    const combined = new Date(date);
    combined.setHours(hours, minutes, 0, 0);
    return combined;
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
}

function formatDateTime(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}
