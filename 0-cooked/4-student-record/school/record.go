package school

// struct StudentRecord which contains a slice of Subjects
type StudentRecord struct {
	// a slice of Subjects
	subjects []Subject
}

// variadit function NewStudentRecord which returns a StudentRecord and accepts one or more Subject in input
// data in ingresso una slice sub di tipo Subject,
// ritorna una collezione di subject (name and score) di nome sr, di tipo StudentRecord
func NewStudentRecord(sub ...Subject) StudentRecord {

	sr := StudentRecord{
		subjects: sub,
	}

	return sr
}

// MaxScoreSubject method is a function which returns the Subject with the highest score
// che prendere come riferimento la struct StudentRecord
func (sr StudentRecord) MaxScoreSubject() Subject {

	max := Subject{
		Name:  "",
		Score: 0,
	}
	// range è un costrutto che sta dentro al for per scorrere una slice
	// in questo caso dobbiamo scorrere il riferimento sr che però non è tipo slice, ma la struct StudentRecord al suo interno ha una slice, quindi basterà accedere a quel campo (sr.subjects)
	for _, subject := range sr.subjects {
		if subject.Score > max.Score {
			max = subject
		}
	}

	return max
}

// method Mean returns the mean of the student
func (sr StudentRecord) Mean() float64 {

	sumScore := 0.0
	for _, subject := range sr.subjects {
		sumScore += subject.Score
	}

	len := float64(len(sr.subjects))

	avg := 0.0
	avg = sumScore / len
	return avg
}
