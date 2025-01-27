// Ref: https://angular-material.dev/articles/angular-material-3

import {
  ChangeDetectorRef,
  Component,
  inject,
  AfterViewInit,
} from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { AsyncPipe, CommonModule } from '@angular/common';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { AuthService } from '../services/auth/auth.service';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { ThemeService } from '../services/theme/theme.service';
import { Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrl: './layout.component.scss',
  imports: [
    CommonModule,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    AsyncPipe,
    RouterLink,
    RouterLinkActive,
    MatSnackBarModule,
  ],
})
export class LayoutComponent implements AfterViewInit {
  private breakpointObserver = inject(BreakpointObserver);
  private authService = inject(AuthService);
  private themeService = inject(ThemeService);
  private router = inject(Router);

  userRole = '';
  username = '';
  isHomePage = false;

  constructor(private cdr: ChangeDetectorRef) {
    this.userRole = this.authService.getUserRole() ?? '';
    this.username = this.authService.getUsername() ?? '';
    this.themeService.initTheme();
    this.router.events
      .pipe(
        filter(
          (event): event is NavigationEnd => event instanceof NavigationEnd,
        ),
      )
      .subscribe((event: NavigationEnd) => {
        this.isHomePage = event.urlAfterRedirects === '/home';
      });
  }

  ngAfterViewInit() {
    this.isHandset$.subscribe(() => {
      this.cdr.detectChanges();
    });
  }

  visibleForRole(role: string): boolean {
    return this.userRole === role;
  }

  isLoggedIn(): boolean {
    if (!this.authService.isLoggedIn()) return false;
    const hasUsername = !!this.username;
    const hasRole = !!this.userRole;
    return hasUsername && hasRole;
  }

  toggleTheme(): void {
    this.themeService.toggleTheme();
  }

  isTheme(theme: 'light' | 'dark'): boolean {
    return this.themeService.getTheme() === theme;
  }

  isHandset$: Observable<boolean> = this.breakpointObserver
    .observe(Breakpoints.Handset)
    .pipe(
      map((result) => result.matches),
      shareReplay(),
    );
}
