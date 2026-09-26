package journal

func (s *FileStore) restoreTail() error {
	records, err := s.readRecords()
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return nil
	}
	last := records[len(records)-1]
	s.sequence = last.Sequence
	s.checksum = last.Checksum
	return nil
}
