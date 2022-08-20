package sec7_2

import "io"

type limitReader struct {
	reader io.Reader
	read   int64
	limit  int64
}

func (r *limitReader) Read(p []byte) (int, error) {
	if int64(len(p)) > r.limit {
		p = p[:r.limit]
	}

	n, err := r.reader.Read(p)
	r.read += int64(n)
	if err != nil {
		return n, err
	}
	if r.read >= r.limit {
		return n, io.EOF
	}
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		reader: r,
		read:   0,
		limit:  n,
	}
}
