import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../services/festival/festival.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { MatTableModule } from '@angular/material/table';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { ItemService } from '../../../services/festival/item.service';
import { UserListResponse } from '../../../models/user/user-responses';
import { UserService } from '../../../services/user/user.service';

@Component({
  selector: 'app-all-accounts',
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
    MatTableModule,
  ],
  templateUrl: './all-accounts.component.html',
  styleUrls: ['./all-accounts.component.scss', '../../../app.component.scss'],
})
export class AllAccountsComponent {
  organizerCount: number = 0;
  employeeCount: number = 0;
  ogranizerCount: number = 0;
  adminCount: number = 0;

  displayColumns = ['id', 'username', 'name', 'role'];

  users: UserListResponse[] = [];

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private userService = inject(UserService);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private itemService = inject(ItemService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadUsers();
  }

  loadUsers() {
    this.userService.getAllUsers().subscribe((response) => {
      console.log(response);
      this.users = response;
      this.organizerCount = this.users.filter(
        (user) => user.role === 'organizer',
      ).length;
      this.employeeCount = this.users.filter(
        (user) => user.role === 'employee',
      ).length;
      this.adminCount = this.users.filter(
        (user) => user.role === 'admin',
      ).length;
    });
  }
}
