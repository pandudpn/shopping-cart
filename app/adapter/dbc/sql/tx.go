package sql

func (sd *SqlDb) Rollback() error {
	return nil
}

func (sd *SqlDb) Commit() error {
	return nil
}

func (sd *SqlDb) TxEnd(txFunc func() error) error {
	return nil
}

func (st *SqlTx) Rollback() error {
	return st.DB.Rollback()
}

func (st *SqlTx) Commit() error {
	return st.DB.Commit()
}

func (st *SqlTx) TxEnd(txFunc func() error) error {
	var err error

	defer func() {
		if p := recover(); p != nil {
			st.Rollback()
			panic(p)
		} else if err != nil {
			st.Rollback()
		} else {
			err = st.Commit()
		}
	}()

	err = txFunc()
	return err
}
