import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
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
import { UserListResponse } from '../../../models/user/user-responses';
import { UserService } from '../../../services/user/user.service';
import { FormsModule } from '@angular/forms';
import { RegisterUserComponent } from '../register-user/register-user.component';

@Component({
  selector: 'app-all-users',
  imports: [
    CommonModule,
    FormsModule,
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
  templateUrl: './all-users.component.html',
  styleUrls: ['./all-users.component.scss', '../../../app.component.scss'],
})
export class AllAccountsComponent {
  private router = inject(Router);
  private userService = inject(UserService);
  private dialog = inject(MatDialog);

  attendeeCount: number = 0;
  organizerCount: number = 0;
  employeeCount: number = 0;
  ogranizerCount: number = 0;
  adminCount: number = 0;

  displayColumns = ['id', 'username', 'name', 'role', 'actions'];

  users: UserListResponse[] = [];

  filterOptions: string[] = [
    'All Roles',
    'Administrator',
    'Organizer',
    'Employee',
    'Attendee',
  ];
  selectedChip = 'All Roles';

  ngOnInit() {
    this.loadUsers();
  }

  loadUsers() {
    this.userService.getAllUsers().subscribe((response) => {
      this.users = response;
      this.attendeeCount = this.users.filter(
        (user) => user.role === 'ATTENDEE',
      ).length;
      this.organizerCount = this.users.filter(
        (user) => user.role === 'ORGANIZER',
      ).length;
      this.employeeCount = this.users.filter(
        (user) => user.role === 'EMPLOYEE',
      ).length;
      this.adminCount = this.users.filter(
        (user) => user.role === 'ADMINISTRATOR',
      ).length;
    });
  }

  get filteredUsers(): UserListResponse[] {
    if (this.selectedChip === 'All Roles') {
      return this.users;
    } else {
      return this.users.filter(
        (user) => user.role === this.selectedChip.toUpperCase(),
      );
    }
  }

  onViewClick(userId: number) {
    this.router.navigate([`/admin/users/${userId}`]);
  }

  onRegisterClick() {
    const dialogRef = this.dialog.open(RegisterUserComponent, {
      width: '800px',
      height: 'auto',
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadUsers();
      }
    });
  }
}
