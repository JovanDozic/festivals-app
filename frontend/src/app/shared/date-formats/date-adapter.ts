// custom-date-adapter.ts
import { Injectable } from '@angular/core';
import { NativeDateAdapter } from '@angular/material/core';

@Injectable()
export class CustomDateAdapter extends NativeDateAdapter {
  override getFirstDayOfWeek(): number {
    // 0 - sunday, 1 - monday, etc.
    return 1;
  }

  // parsing to 'dd.MM.yyyy'
  override parse(value: string): Date | null {
    if (typeof value === 'string' && value.length > 0) {
      const parts = value.split('.');
      if (parts.length === 3) {
        const day = Number(parts[0]);
        const month = Number(parts[1]) - 1; // months are 0-based
        const year = Number(parts[2]);
        const date = new Date(year, month, day);
        if (!isNaN(date.getTime())) {
          return date;
        }
      }
    }
    return null;
  }

  override format(date: Date): string {
    const day = this._to2digit(date.getDate());
    const month = this._to2digit(date.getMonth() + 1);
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
  }

  private _to2digit(n: number) {
    return ('00' + n).slice(-2);
  }
}
