import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import { ReactiveFormsModule } from '@angular/forms';
import { FestivalService } from '../../../../services/festival/festival.service';
import { Employee } from '../../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { MatTableModule } from '@angular/material/table';
import { MatMenuModule } from '@angular/material/menu';
import { ActivatedRoute, Router } from '@angular/router';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../../shared/confirmation-dialog/confirmation-dialog.component';

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
    '../../../../app.component.scss',
  ],
})
export class AddExistingEmployeeComponent implements OnInit {
  private dialogRef = inject(MatDialogRef<AddExistingEmployeeComponent>);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);

  isLoading = false;
  employeeCount = 0;
  employees: Employee[] = [];
  displayedColumns = ['username', 'email', 'name', 'actions'];
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);

  ngOnInit(): void {
    this.loadEmployees();
  }

  loadEmployees() {
    if (this.data.festivalId) {
      this.festivalService
        .getEmployeesNotOnFestival(Number(this.data.festivalId))
        .subscribe({
          next: (employees) => {
            this.employees = employees;
            this.employeeCount = employees.length;
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error getting employees');
            this.employees = [];
          },
        });
    }
  }

  onAddEmployeeClick(employee: Employee) {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Add an Employee',
        message: `Are you sure you want to add ${employee.firstName} ${employee.lastName} to this Festival?`,
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.addEmployee(employee);
      }
    });
  }

  addEmployee(employee: Employee) {
    if (this.data.festivalId) {
      this.festivalService
        .employEmployee(this.data.festivalId, employee.id)
        .subscribe({
          next: () => {
            this.snackbarService.show(
              `${employee.firstName} ${employee.lastName} added to festival`,
            );
            this.loadEmployees();
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error adding employee');
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
