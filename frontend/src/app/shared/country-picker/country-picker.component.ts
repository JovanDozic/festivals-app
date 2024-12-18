import { Component, OnInit, forwardRef } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { NG_VALUE_ACCESSOR, ControlValueAccessor } from '@angular/forms';
import { Observable, of } from 'rxjs';
import { map, startWith } from 'rxjs/operators';
import * as countries from 'i18n-iso-countries';
import enLocale from 'i18n-iso-countries/langs/en.json';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule, MatSelectChange } from '@angular/material/select';
import { NgxMatSelectSearchModule } from 'ngx-mat-select-search';
import { CommonModule } from '@angular/common';
import { MatOptionModule } from '@angular/material/core';

interface Country {
  name: string;
  iso3: string;
}

@Component({
  selector: 'app-country-picker',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatSelectModule,
    NgxMatSelectSearchModule,
    MatOptionModule,
  ],
  templateUrl: './country-picker.component.html',
  styleUrls: ['./country-picker.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => CountryPickerComponent),
      multi: true,
    },
  ],
})
export class CountryPickerComponent implements ControlValueAccessor {
  countryFilterCtrl: FormControl = new FormControl('');
  countriesList: Country[] = [];
  filteredCountries$: Observable<Country[]>;
  selectedCountryISO3: string | null = null;

  private onChange: (value: string | null) => void = () => {};
  private onTouched: () => void = () => {};

  constructor() {
    countries.registerLocale(enLocale);

    this.countriesList = Object.entries(
      countries.getNames('en', { select: 'official' }),
    )
      .map(([iso2, name]) => {
        const iso3 = countries.alpha2ToAlpha3(iso2);
        if (iso3) {
          return { name, iso3 };
        }
        return null;
      })
      .filter((country): country is Country => country !== null);

    this.filteredCountries$ = this.countryFilterCtrl.valueChanges.pipe(
      startWith(''),
      map((value) => this._filterCountries(value || '')),
    );
  }

  private _filterCountries(value: string): Country[] {
    const filterValue = value.toLowerCase();
    return this.countriesList.filter((country) =>
      country.name.toLowerCase().includes(filterValue),
    );
  }

  // ControlValueAccessor Methods
  writeValue(value: string | null): void {
    this.selectedCountryISO3 = value;
  }

  registerOnChange(fn: (value: string | null) => void): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void): void {
    this.onTouched = fn;
  }

  setDisabledState?(isDisabled: boolean): void {
    if (isDisabled) {
      this.countryFilterCtrl.disable();
    } else {
      this.countryFilterCtrl.enable();
    }
  }

  onSelectionChange(event: MatSelectChange) {
    this.selectedCountryISO3 = event.value;
    this.onChange(this.selectedCountryISO3);
    this.onTouched();
  }

  // Optional: Comparison function if needed
  compareIso3(c1: string, c2: string): boolean {
    return c1 === c2;
  }
}
