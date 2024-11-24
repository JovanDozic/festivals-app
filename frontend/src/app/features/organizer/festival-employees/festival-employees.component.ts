import { Component, inject, OnInit } from '@angular/core';
import { Employee, Festival } from '../../../models/festival/festival.model';
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
import {
  ConfirmationDialog,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { EditFestivalComponent } from '../edit-festival/edit-festival.component';
import { animate, style, transition, trigger } from '@angular/animations';
import { MatMenuModule } from '@angular/material/menu';
import { MatTableModule } from '@angular/material/table';

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
    '../../../app.component.scss',
  ],
})
export class FestivalEmployeesComponent implements OnInit {
  onRegisterEmployee() {
    throw new Error('Method not implemented.');
  }
  onAddEmployeeClick() {
    throw new Error('Method not implemented.');
  }
  festival: Festival | null = null;
  isLoading = false;
  employeeCount: number = 0;
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
          console.log('Festival: ', festival);
          this.festival = festival;
          this.isLoading = false;
        },
        error: (error) => {
          console.log('Error fetching festival information: ', error);
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
          console.log('Employee count: ', count);
          this.employeeCount = count;
        },
        error: (error) => {
          console.log('Error fetching employee count: ', error);
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
}
