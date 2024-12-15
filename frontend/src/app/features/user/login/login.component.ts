import { Component, Inject, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { AuthService } from '../../../services/auth/auth.service';
import { MatCardModule } from '@angular/material/card';
import { CommonModule } from '@angular/common';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss', '../../../app.component.scss'],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
  ],
})
export class LoginComponent {
  private fb = inject(FormBuilder);
  private authService = inject(AuthService);
  private router = inject(Router);
  private snackbarService = inject(SnackbarService);

  loginForm = this.fb.group({
    username: ['', [Validators.required]],
    password: ['', [Validators.required]],
  });

  errorMessage: string | null = null;

  login() {
    if (this.loginForm.valid) {
      this.authService
        .login(this.loginForm.value as { username: string; password: string })
        .subscribe({
          error: (err) => {
            console.error('Login error:', err);
            this.snackbarService.show('Invalid username or password');
          },
        });
    }
  }

  navigateToRegister() {
    this.router.navigate(['/register']);
  }
}
