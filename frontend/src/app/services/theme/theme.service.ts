import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ThemeService {
  private readonly THEME_KEY = 'selected-theme';

  setTheme(theme: 'light' | 'dark'): void {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem(this.THEME_KEY, theme);
  }

  getTheme(): 'light' | 'dark' {
    return (
      (localStorage.getItem(this.THEME_KEY) as 'light' | 'dark') || 'light'
    );
  }

  toggleTheme(): void {
    const currentTheme = this.getTheme();
    const newTheme = currentTheme === 'light' ? 'dark' : 'light';
    this.setTheme(newTheme);
  }

  initTheme(): void {
    const savedTheme = this.getTheme();
    this.setTheme(savedTheme);
  }
}
