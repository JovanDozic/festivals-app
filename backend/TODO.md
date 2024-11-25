can i get a list of employees that work on any festival that i'm organizer of?

--

[x] create employee: POST organizer/employee

- ovo ce da se kreira novi nalog i profil
- onda ti treba i endpoint call da ga zaposlis

[] get all employees for my organization: GET organizer/employee

- this will include festival IDs that they're working on so organizer can change it

[x] get employees on the festival: GET organizer/festival/{festivalId}/employee

[] get employee: GET organizer/employee/{employeeId}

[] update employee: PUT organizer/employee/{employeeId}

[] delete employee: DELETE organizer/{employeeId}

[x] add employee to the festival: PUT organizer/festival/{festivalId}/employee/{employeeId}/employ

[] remove employee from the festival: PUT organizer/festival/{festivalId}/employee/{employeeId}/fire

[] number of employees in the festival: GET organizer/festival/{festivalId}/employee/count

[] GET all available employees that can be added to the festival

- should it be all employees in the system? or employees that are not in any festival?
- we should list out all employees that are not in any festival
- and also list should not include any employee that's already working on that festival
- so endpoint should look like
  GET organizer/festival/{festivalId}/employee/available
- so we have festival ID against which we're checking if employee is available
