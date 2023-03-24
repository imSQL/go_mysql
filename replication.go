package imSQL

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/juju/errors"
)

type (

	// Slave change master to
	Slave struct {
		MasterBind                string   `json:"master_bind" db:"master_bind"`
		MasterHost                string   `json:"master_host" db:"master_host"`
		MasterUser                string   `json:"master_user" db:"master_user"`
		MasterPassword            string   `json:"master_password" db:"master_password"`
		MasterPort                uint64   `json:"master_port" db:"master_port"`
		MasterConnectRetry        uint64   `json:"master_connect_retry" db:"master_connect_retry"`
		MasterRetryCount          uint64   `json:"master_retry_count" db:"master_retry_count"`
		MasterDelay               uint64   `json:"master_delay" db"master_delay"`
		MasterHeartbeatPeriod     uint64   `json:"master_heartbeat_period" db:"master_heartbeat_period"`
		MasterLogFile             string   `json:"master_log_file" db:"master_log_file"`
		MasterLogPos              uint64   `json:"master_log_pos" db:"master_log_pos"`
		MasterAutoPosition        uint64   `json:"master_auto_position" db:"master_auto_position"`
		RelayLogFile              string   `json:"relay_log_file" db:"relay_log_file"`
		RelayLogPos               uint64   `json:"relay_log_pos" db:"relay_log_pos"`
		MasterSSL                 uint64   `json:"master_ssl" db:"master_ssl"`
		MasterSSLCa               string   `json:"master_ssl_ca" db:"master_ssl_ca"`
		MasterSSLCapath           string   `json:"master_ssl_capath" db:"master_ssl_capath"`
		MasterSSLCert             string   `json:"master_ssl_cert" db:"master_ssl_cert"`
		MasterSSLCrl              string   `json:"master_ssl_crl" db:"master_ssl_crl"`
		MasterSSLCrlpath          string   `json:"master_ssl_crlpath" db:"master_ssl_crlpath"`
		MasterSSLKey              string   `json:"master_ssl_key" db:"master_ssl_key"`
		MasterSSLCipher           string   `json:"master_ssl_cipher" db:"master_ssl_cipher"`
		MasterSSLVerifyServerCert uint64   `json:"master_ssl_verify_server_cert" db:"master_ssl_verify_server_cert"`
		MasterTLSVersion          string   `json:"master_tls_version" db:"master_tls_version"`
		IgnoreServerIds           []string `json:"ignore_server_ids" db:"ignore_server_ids"`
	}
)

// NewSlave new slave struct.
func NewSlave(host string, port uint64, username string, password string, filename string, filepos uint64) (*Slave, error) {

	slave := new(Slave)

	slave.MasterHost = host
	slave.MasterPort = port
	slave.MasterUser = username
	slave.MasterPassword = password
	slave.MasterLogFile = filename
	slave.MasterLogPos = filepos

	slave.MasterBind = ""
	slave.MasterConnectRetry = 60
	slave.MasterRetryCount = 86400
	slave.MasterDelay = 0
	slave.MasterHeartbeatPeriod = 0
	slave.MasterAutoPosition = 0

	slave.RelayLogFile = ""
	slave.RelayLogPos = 0

	slave.MasterSSL = 0
	slave.MasterSSLCa = ""
	slave.MasterSSLCapath = ""
	slave.MasterSSLCapath = ""
	slave.MasterSSLCert = ""
	slave.MasterSSLCrl = ""
	slave.MasterSSLCrlpath = ""
	slave.MasterSSLKey = ""
	slave.MasterSSLCipher = ""
	slave.MasterSSLVerifyServerCert = 0
	slave.MasterTLSVersion = ""
	slave.IgnoreServerIds = []string{}

	return slave, nil
}

// SetMasterBind set master bind
func (slave *Slave) SetMasterBind(master_bind string) {
	slave.MasterBind = master_bind
}

// SetMasterConnectRetry set master connect retry
func (slave *Slave) SetMasterConnectRetry(master_connect_retry uint64) {
	slave.MasterConnectRetry = master_connect_retry
}

// SetMasterRetryCount set master retry count
func (slave *Slave) SetMasterRetryCount(master_retry_count uint64) {
	slave.MasterRetryCount = master_retry_count
}

// SetMasterDelay set master delay
func (slave *Slave) SetMasterDelay(master_delay uint64) {
	slave.MasterDelay = master_delay
}

// SetMasterHeartbeatPeriod set master heartbeat period
func (slave *Slave) SetMasterHeartbeatPeriod(master_heartbeat_period uint64) {
	slave.MasterHeartbeatPeriod = master_heartbeat_period
}

// SetMasterAutoPosition set master auto position
func (slave *Slave) SetMasterAutoPosition(auto bool) {
	if auto {
		slave.MasterAutoPosition = 1
	} else {
		slave.MasterAutoPosition = 0
	}
}

// SetRelayLogFile set relay log file
func (slave *Slave) SetRelayLogFile(relay_log_file string) {
	slave.RelayLogFile = relay_log_file
}

// SetRelayLogPos set relay log pos
func (slave *Slave) SetRelayLogPos(relay_log_pos uint64) {
	slave.RelayLogPos = relay_log_pos
}

// SetMasterSSL set master ssl
func (slave *Slave) SetMasterSSL(ssl bool) {
	if ssl {
		slave.MasterSSL = 1
	} else {
		slave.MasterSSL = 0
	}
}

// SetMasterSSLCa set master ssl ca
func (slave *Slave) SetMasterSSLCa(master_ca string) {
	slave.MasterSSLCa = master_ca
}

// SetMasterSSLCaPath set master ssl capath
func (slave *Slave) SetMasterSSLCaPath(capath string) {
	slave.MasterSSLCapath = capath
}

// SetMasterSSLCert set master ssl cert
func (slave *Slave) SetMasterSSLCert(cert string) {
	slave.MasterSSLCert = cert
}

// SetMasterSSLCrl set master ssl crl
func (slave *Slave) SetMasterSSLCrl(crl string) {
	slave.MasterSSLCrl = crl
}

// SetMasterSSLCrlPath set master ssl crlpath
func (slave *Slave) SetMasterSSLCrlPath(crlpath string) {
	slave.MasterSSLCrlpath = crlpath
}

// SetMasterSSLKey set master ssl key
func (slave *Slave) SetMasterSSLKey(key string) {
	slave.MasterSSLKey = key
}

// SetMasterSSLCipher set master ssl cipher
func (slave *Slave) SetMasterSSLCipher(cipher string) {
	slave.MasterSSLCipher = cipher
}

// SetMasterSSLVerifyServerCert set master ssl verify server cert
func (slave *Slave) SetMasterSSLVerifyServerCert(verify bool) {
	if verify {
		slave.MasterSSLVerifyServerCert = 1
	} else {
		slave.MasterSSLVerifyServerCert = 0
	}
}

// SetMasterTlsVersion set master tls version
func (slave *Slave) SetMasterTlsVersion(version string) {
	slave.MasterTLSVersion = version
}

// SetIgnoreIds set ignore ids
func (slave *Slave) SetIgnoreIds(ids ...string) {
	slave.IgnoreServerIds = append(slave.IgnoreServerIds, ids...)
}

func (slave *Slave) ChangeMaster(db *sql.DB) error {

	// define args
	args := make([]string, 0, 50)

	args = append(args, fmt.Sprintf("MASTER_HOST='%s'", slave.MasterHost))
	args = append(args, fmt.Sprintf("MASTER_PORT=%d", slave.MasterPort))
	args = append(args, fmt.Sprintf("MASTER_USER='%s'", slave.MasterUser))
	args = append(args, fmt.Sprintf("MASTER_PASSWORD='%s'", slave.MasterPassword))
	args = append(args, fmt.Sprintf("MASTER_LOG_FILE='%s'", slave.MasterLogFile))
	args = append(args, fmt.Sprintf("MASTER_LOG_POS=%d", slave.MasterLogPos))

	args = append(args, fmt.Sprintf("MASTER_BIND='%s'", slave.MasterBind))

	if slave.MasterConnectRetry != 0 {
		args = append(args, fmt.Sprintf("MASTER_CONNECT_RETR=%d", slave.MasterConnectRetry))
	}

	if slave.MasterRetryCount != 0 {
		args = append(args, fmt.Sprintf("MASTER_RETRY_COUNT=%d", slave.MasterRetryCount))
	}

	if slave.MasterDelay != 0 {
		args = append(args, fmt.Sprintf("MASTER_DELAY=%d", slave.MasterDelay))
	}

	if slave.MasterHeartbeatPeriod != 0 {
		args = append(args, fmt.Sprintf("MASTER_HEARTBEAT_PERIOD=%d", slave.MasterHeartbeatPeriod))
	}

	if slave.MasterAutoPosition == 1 {
		args = append(args, fmt.Sprintf("MASTER_AUTO_POSITION=1"))
	}

	if len(slave.RelayLogFile) != 0 {
		args = append(args, fmt.Sprintf("RELAY_LOG_FILE='%s'", slave.RelayLogFile))
	}

	if slave.RelayLogPos != 0 {
		args = append(args, fmt.Sprintf("RELAY_LOG_POS=%d", slave.RelayLogPos))
	}

	if slave.MasterSSL == 1 {
		args = append(args, fmt.Sprintf("MASTER_SSL=1"))
	}

	if len(slave.MasterSSLCa) != 0 {
		args = append(args, fmt.Sprintf("MASTER_SSL_CA='%s'", slave.MasterSSLCa))
	}

	if len(slave.MasterSSLCapath) != 0 {
		args = append(args, fmt.Sprintf("MASTER_SSL_CAPATH='%s'", slave.MasterSSLCapath))
	}

	if len(slave.MasterSSLCert) != 0 {
		args = append(args, fmt.Sprintf("MASTER_SSL_CERT='%s'", slave.MasterSSLCert))
	}

	if len(slave.MasterSSLCrl) != 0 {
		args = append(args, fmt.Sprintf("MASTER_SSL_CRL='%s'", slave.MasterSSLCert))
	}

	//if len(slave.MasterSSLCrlpath) != 0 {
	//	args = append(args, fmt.Sprintf("MASTER_SSL_KEY='%s'", slave.SetMasterSSLKey))
	//}

	//if len(slave.MasterSSLKey) != 0 {
	//	args = append(args, fmt.Sprintf("MASTER_SSL_KEY='%s'", slave.SetMasterSSLKey))
	//}

	if len(slave.MasterSSLCipher) != 0 {
		args = append(args, fmt.Sprintf("MASTER_SSL_CIPHER='%s'", slave.MasterSSLCipher))
	}

	if slave.MasterSSLVerifyServerCert != 0 {
		args = append(args, fmt.Sprintf("MASTER_SSL_VERIFY_SERVER_CERT=%d", slave.MasterSSLVerifyServerCert))
	}

	if len(slave.MasterTLSVersion) != 0 {
		args = append(args, fmt.Sprintf("MASTER_TLS_VERSION='%s'", slave.MasterTLSVersion))
	}

	if len(slave.IgnoreServerIds) != 0 {
		args = append(args, fmt.Sprintf("IGNORE_SERVER_IDS='%s'", strings.Join(slave.IgnoreServerIds, ",")))
	}

	ChangeMaster := fmt.Sprintf("CHANGE MASTER TO %s", strings.Join(args, ","))

	_, err := db.Exec(ChangeMaster)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
