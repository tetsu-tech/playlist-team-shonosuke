// not Found.tsxに書き方を乗っ取る
export function login() {
    return 
    <div className="input">
        <div className="title">
            <div className="title_text">ログイン</div>
        </div>
        <div className="large_background">
            <div className="login_info_input">
                <div id="id2" className="id_input">
                    <div className="id_tag">
                        <div className="id_title">ID</div>
                    </div>
                    <input className="id_text" type="text" id="id_text" name="id_text"></input>
                </div>
                <div id="id2" className="pass_input">
                    <div className="pass_tag">
                        <div className="pass_title">PASSWORD</div>
                    </div>
                    <input className="pass_text" type="text" id="pass_text" name="pass_text"></input>
                </div>
            </div>
            <div className="button_frame">
                <input className="button" type="button" value="ログインする"></input>
            </div>
        </div>
    </div>
    ;
  }