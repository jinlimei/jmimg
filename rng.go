package jmimg

type UploadNameGenerator func(input string) string

func (s *Service) SetUploadNameGenerator(g UploadNameGenerator) {
	s.nameGenerator = g
}
