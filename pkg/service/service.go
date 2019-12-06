package v1

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/laughtt/loginService/api/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

//AuthServiceServer connection
type AuthServiceServer struct {
	db *sql.DB
}

//NewAuthServiceServer create a AUTHservice
func NewAuthServiceServer(dd *sql.DB) v1.AuthServiceServer {
	return &AuthServiceServer{db: dd}
}

//Config  datbase connection
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
	// DB Datastore parameters selsction
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

//Chekea si el api es la version correcta
func (s *AuthServiceServer) checkAPI(api string) error {

	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}


// func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
//     tx, err := db.Begin()
//     if err != nil {
//         return
//     }
//     defer func() {
//         if p := recover(); p != nil {
//             tx.Rollback()
//             panic(p) // re-throw panic after Rollback
//         } else if err != nil {
//             tx.Rollback() // err is non-nil; don't change it
//         } else {
//             err = tx.Commit() // err is nil; if Commit returns error update err
//         }
//     }()
//     err = txFunc(tx)
// 	return err
// 
// }

// func (s Service) DoSomething() error {
//     return Transact(s.db, func (tx *sql.Tx) error {
//         if _, err := tx.Exec(...); err != nil {
//             return err
//         }
//         if _, err := tx.Exec(...); err != nil {
//             return err
//         }
//         return nil
//     })
// }

//Connect database
func (s *AuthServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
func (s *AuthServiceServer) beginConnection(ctx context.Context) (*sql.Tx , error){
	b , err := s.db.Begin()
	
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return b, nil
}

//CreateAccount a new user for the database
func (s *AuthServiceServer) CreateAccount(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c, err := s.connect(ctx)

	if err != nil {
		return nil, err
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(req.Data.GetReminder())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}
	password , username := req.Data.GetPassword() , req.Data.GetEmail()
	res, err := c.ExecContext(ctx, "INSERT INTO members(`password`, `username`, `Reminder`) VALUES(?, ?, ?)", password, username , reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into members-> "+err.Error())
	}
	_ , err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
	}
	res , err = c.ExecContext(ctx, "INSERT INTO logs(`action`, `reminder`, `user`) VALUES(? , ? , ? )", "CREATE", reminder , username )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into logs-> "+err.Error())
	}
	_ , err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created logs-> "+err.Error())
	}
	return &v1.CreateResponse{
		Api: apiVersion,
		Success: true,
		Error: "None",
		}, nil
}

//EraseAccount an user from the database
func (s *AuthServiceServer) EraseAccount(ctx context.Context, req *v1.EraseAccountRequest) (*v1.EraseAccountResponse, error) {
	return &v1.EraseAccountResponse{}, nil
}

//LoginAccount an account in the server
func (s *AuthServiceServer) LoginAccount(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return &v1.LoginResponse{}, nil
}

//ChangePassword the password from a user
func (s *AuthServiceServer) ChangePassword(ctx context.Context, req *v1.ChangePasswordRequest) (*v1.ChangePasswordResponse, error) {
	return &v1.ChangePasswordResponse{}, nil
}