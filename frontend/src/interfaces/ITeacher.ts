import { EducationalInterface } from "./IEducational";
import { FacultyInterface } from "./IFaculty";
import { OfficerInterface} from "./IOfficer"
import { PrefixInterface } from "./IPrefix";


export interface TeacherInterface {
    ID?: number,
    Name?: string;
    Email?:         string ;
	Password?:	string;
	Officer_ID ?:    number;
	Officer ?: OfficerInterface;
	Faculty_ID ?:    number;
	Faculty ?: FacultyInterface;
	Prefix_ID?:    number;
	Prefix?: PrefixInterface;
	Educational_ID?:    number;
	Educational?: EducationalInterface;

   }