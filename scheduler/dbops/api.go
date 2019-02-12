package dbops

import "log"

// 1. user->api service->delete video
// 2. api service->scheduler->write video deletion record
// 3. timer
// 4. timer->runner->read wvdr->exec->delete video from folder

func AddVideoDeletionRecord(vid string) error {
	stmrIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) values (?)")
	if err != nil {
		log.Printf("AddVideoDeletionRecord %v", err)
		return err
	}

	_, err = stmrIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord %v", err)
		return err
	}

	defer stmrIns.Close()
	return nil

}
