// custom-date-adapter.ts
import { Injectable } from '@angular/core';
import { NativeDateAdapter } from '@angular/material/core';

@Injectable()
export class CustomDateAdapter extends NativeDateAdapter {
  // Override to set Monday as the first day of the week
  override getFirstDayOfWeek(): number {
    return 1; // 0 = Sunday, 1 = Monday, etc.
  }

  // Override the parse method to handle 'dd.MM.yyyy' format
  override parse(value: any): Date | null {
    if (typeof value === 'string' && value.length > 0) {
      const parts = value.split('.');
      if (parts.length === 3) {
        const day = Number(parts[0]);
        const month = Number(parts[1]) - 1; // Months are 0-based
        const year = Number(parts[2]);
        const date = new Date(year, month, day);
        if (!isNaN(date.getTime())) {
          return date;
        }
      }
    }
    return null;
  }

  // Optionally, override the format method if you want to ensure consistency
  override format(date: Date, displayFormat: Object): string {
    const day = this._to2digit(date.getDate());
    const month = this._to2digit(date.getMonth() + 1);
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
  }

  private _to2digit(n: number) {
    return ('00' + n).slice(-2);
  }
}
