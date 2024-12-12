import { Component, inject, OnInit } from '@angular/core';
import { AuthService } from '../../../services/auth/auth.service';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { UserService } from '../../../services/user/user.service';
import { UserProfileResponse } from '../../../models/user/user-responses';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { CommonModule } from '@angular/common';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { MatTooltipModule } from '@angular/material/tooltip';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-user',
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
  ],
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss', '../../../app.component.scss'],
})
export class AccountComponent {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private authService = inject(AuthService);
  private userService = inject(UserService);
  readonly dialog = inject(MatDialog);

  userProfile: UserProfileResponse | null = null;

  ngOnInit() {
    this.getUserProfile();
  }

  goBack() {
    this.router.navigate([`/admin/users`]);
  }

  getUserProfile() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.userService.getUserById(Number(id)).subscribe((response) => {
        this.userProfile = response;
      });
    }
  }
}
