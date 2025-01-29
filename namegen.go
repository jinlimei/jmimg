package jmimg

type UploadNameGenerator func(input string, ext string) string

func (s *Service) SetUploadNameGenerator(g UploadNameGenerator) {
	s.nameGenerator = g
}
