import { Component, inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDatepickerModule } from '@angular/material/datepicker';
import {
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
} from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { UserService } from '../../../services/user/user.service';
import { provideNativeDateAdapter } from '@angular/material/core';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';

@Component({
  selector: 'app-edit-festival',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatInputModule,
    MatDatepickerModule,
  ],
  templateUrl: './edit-festival.component.html',
  // todo: remove app styles maybe idk what are they doing
  styleUrls: ['./edit-festival.component.scss', '../../../app.component.scss'],
  providers: [provideNativeDateAdapter()],
})
export class EditFestivalComponent implements OnInit {
  readonly dialogRef = inject(MatDialogRef<EditFestivalComponent>);
  readonly formBuilder = inject(FormBuilder);
  readonly data = inject<any>(MAT_DIALOG_DATA);
  private snackbarService = inject(SnackbarService);

  editFestivalForm: FormGroup = this.formBuilder.group({
    name: ['', Validators.required],
  });

  ngOnInit() {
    if (this.data) {
      const { name } = this.data;
      this.editFestivalForm.patchValue({
        name,
        // dateOfBirth: dateOfBirth ? new Date(dateOfBirth) : null,
        // phoneNumber,
      });
    }
  }

  editFestival() {
    if (this.editFestivalForm.valid) {
      // todo: call service
      this.dialogRef.close(true);
      this.snackbarService.show('Festival updated successfully');
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }
}
