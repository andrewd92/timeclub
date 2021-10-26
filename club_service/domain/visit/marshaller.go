package visit

type visitJson struct {
	Id         int64  `json:"id"`
	Start      string `json:"start"`
	ClientId   int64  `json:"client_id"`
	ClientName string `json:"client_name"`
	ClubId     int64  `json:"club_id"`
	Comment    string `json:"comment"`
}

func (v Visit) Marshal() interface{} {
	return visitJson{
		Id:         v.id,
		Start:      v.start.Format("2006-01-02 15:04"),
		ClientId:   v.client.Id(),
		ClientName: v.client.Name(),
		ClubId:     v.club.Id(),
		Comment:    v.comment,
	}
}

func MarshalAll(visits []*Visit) []interface{} {
	result := make([]interface{}, len(visits))

	for i, visit := range visits {
		result[i] = visit.Marshal()
	}

	return result
}
