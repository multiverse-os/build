package build

type Developer struct {
	Name    string
	Email   string
	Website string
	Keys    []Keypair
}

// TODO:
// AddDeveloper, SaveEdit, UpdateProfile, AssignTask, UpdateTask,

func AddKey(key KeyType, publicKey PublicKey) error {
	return nil
}
