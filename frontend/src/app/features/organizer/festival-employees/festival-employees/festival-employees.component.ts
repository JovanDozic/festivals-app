import { Component, inject, OnInit } from '@angular/core';
import { Employee, Festival } from '../../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../../services/festival/festival.service';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
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
import { RegisterEmployeeComponent } from '../register-employee/register-employee.component';
import { AddExistingEmployeeComponent } from '../add-existing-employee/add-existing-employee.component';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../../shared/confirmation-dialog/confirmation-dialog.component';
import { EditEmployeeComponent } from '../edit-employee/edit-employee.component';

@Component({
  selector: 'app-festival-employees',
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
  templateUrl: './festival-employees.component.html',
  styleUrls: [
    './festival-employees.component.scss',
    '../../../../app.component.scss',
  ],
})
export class FestivalEmployeesComponent implements OnInit {
  festival: Festival | null = null;
  isLoading = false;
  employeeCount = 0;
  employees: Employee[] = [];
  displayedColumns = ['username', 'email', 'name', 'phoneNumber', 'actions'];

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialog = inject(MatDialog);

  ngOnInit() {
    this.loadFestival();
    this.loadEmployeeCount();
    this.loadEmployees();
  }

  goBack() {
    this.router.navigate([`organizer/my-festivals/${this.festival?.id}`]);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
          this.isLoading = true;
        },
      });
    }
  }

  loadEmployeeCount() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getEmployeeCount(Number(id)).subscribe({
        next: (count) => {
          this.employeeCount = count;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting employee count');
          this.employeeCount = 0;
        },
      });
    }
  }

  loadEmployees() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getEmployees(Number(id)).subscribe({
        next: (employees) => {
          this.employees = employees;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting employees');
          this.employees = [];
        },
      });
    }
  }

  onEditEmployeeClick(employee: Employee) {
    const dialogRef = this.dialog.open(EditEmployeeComponent, {
      data: employee,
      width: '800px',
      height: 'auto',
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadEmployeeCount();
        this.loadEmployees();
      }
    });
  }

  onAddEmployeeClick() {
    const dialogRef = this.dialog.open(AddExistingEmployeeComponent, {
      data: {
        festivalId: this.festival?.id,
      },
      width: '800px',
      height: 'auto',
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadEmployeeCount();
        this.loadEmployees();
      }
    });
  }

  onRegisterEmployee() {
    const dialogRef = this.dialog.open(RegisterEmployeeComponent, {
      data: {
        festivalId: this.festival?.id,
      },
      width: '800px',
      height: 'auto',
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadEmployeeCount();
        this.loadEmployees();
      }
    });
  }

  onFireEmployeeClick(employee: Employee) {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Fire Employee',
        message: `Are you sure you want to fire ${employee.firstName} ${employee.lastName}?`,
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.fireEmployee(employee);
      }
    });
  }

  fireEmployee(employee: Employee) {
    if (this.festival) {
      this.festivalService
        .fireEmployee(this.festival.id, employee.id)
        .subscribe({
          next: () => {
            this.snackbarService.show(
              `${employee.firstName} ${employee.lastName} removed from the festival`,
            );
            this.loadEmployeeCount();
            this.loadEmployees();
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error firing employee');
          },
        });
    }
  }
}
