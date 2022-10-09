// author: ashing
// time: 2020/4/16 11:39 下午
// mail: axingfly@gmail.com
// Less is more.

package dictionary

type Dictionary map[string]string

//var (
//	ErrNotFound   = errors.New("could not find the word you were looking for")
//	ErrWordExists = errors.New("cannot add word because it already exists")
//)

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	//return ""
	var (
		value string
		ok    bool
	)
	if value, ok = d[key]; ok {
		return value, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
	//_, ok := d[key]
	//if !ok {
	//	d[key] = value
	//} else {
	//	return ErrWordExists
	//}
	//return nil
}

func (d Dictionary) Update(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil

	//_, ok := d[key]
	//if !ok {
	//	return ErrWordDoesNotExist
	//}
	//d[key] = value
	//return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
