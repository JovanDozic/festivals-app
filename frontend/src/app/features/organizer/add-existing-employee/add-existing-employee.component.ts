import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { FestivalService } from '../../../services/festival/festival.service';
import {
  CreateStaffRequest,
  Employee,
} from '../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../services/user/user.service';
import { CreateProfileRequest } from '../../../models/user/user-profile-request.model';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { MatTableModule } from '@angular/material/table';
import { MatMenuModule } from '@angular/material/menu';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-add-existing-employee',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatGridListModule,
    MatIconModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatTabsModule,
    MatTableModule,
    MatMenuModule,
  ],
  templateUrl: './add-existing-employee.component.html',
  styleUrls: [
    './add-existing-employee.component.scss',
    '../../../app.component.scss',
  ],
})
export class AddExistingEmployeeComponent implements OnInit {
  private dialogRef = inject(MatDialogRef<AddExistingEmployeeComponent>);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);

  isLoading = false;
  employeeCount: number = 0;
  employees: Employee[] = [];
  displayedColumns = ['username', 'email', 'name', 'actions'];
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);

  ngOnInit(): void {
    console.log('Festival ID: ', this.data.festivalId);
    this.loadEmployees();
  }

  loadEmployees() {
    if (this.data.festivalId) {
      this.festivalService
        .getEmployeesNotOnFestival(Number(this.data.festivalId))
        .subscribe({
          next: (employees) => {
            console.log('Employees: ', employees);
            this.employees = employees;
          },
          error: (error) => {
            console.log('Error fetching employees: ', error);
            this.snackbarService.show('Error getting employees');
            this.employees = [];
          },
        });
    }
  }

  done() {
    this.dialogRef.close(true);
  }
  closeDialog() {
    this.dialogRef.close(false);
  }
}
