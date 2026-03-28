#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
# shellcheck disable=SC1091
source "$SCRIPT_DIR/accounts.env"

register_user() {
  local email="$1"
  local username="$2"
  local password="$3"
  local first_name="$4"
  local last_name="$5"

  curl -s -X POST "$API_BASE_URL/api/auth/register" \
    -H "Content-Type: application/json" \
    -d "{\"email\":\"$email\",\"username\":\"$username\",\"password\":\"$password\",\"firstName\":\"$first_name\",\"lastName\":\"$last_name\"}" >/dev/null || true
}

login_and_extract_token() {
  local email="$1"
  local password="$2"

  curl -s -X POST "$API_BASE_URL/api/auth/login" \
    -H "Content-Type: application/json" \
    -d "{\"email\":\"$email\",\"password\":\"$password\"}" | \
    node -e 'const fs=require("fs"); const input=fs.readFileSync(0,"utf8"); const data=JSON.parse(input); if(!data.token){ console.error(input); process.exit(1);} process.stdout.write(data.token)'
}

extract_user_id() {
  local token="$1"
  TOKEN="$token" node -e 'const payload=JSON.parse(Buffer.from(process.env.TOKEN.split(".")[1],"base64url").toString()); process.stdout.write(payload.userId)'
}

register_user "$ALICE_EMAIL" "$ALICE_USERNAME" "$ALICE_PASSWORD" "$ALICE_FIRST_NAME" "$ALICE_LAST_NAME"
register_user "$BOB_EMAIL" "$BOB_USERNAME" "$BOB_PASSWORD" "$BOB_FIRST_NAME" "$BOB_LAST_NAME"

ALICE_TOKEN="$(login_and_extract_token "$ALICE_EMAIL" "$ALICE_PASSWORD")"
BOB_TOKEN="$(login_and_extract_token "$BOB_EMAIL" "$BOB_PASSWORD")"
ALICE_USER_ID="$(extract_user_id "$ALICE_TOKEN")"
BOB_USER_ID="$(extract_user_id "$BOB_TOKEN")"

cat > "$SCRIPT_DIR/dm_session.env" <<EOF
export API_BASE_URL='$API_BASE_URL'
export ALICE_EMAIL='$ALICE_EMAIL'
export ALICE_PASSWORD='$ALICE_PASSWORD'
export ALICE_TOKEN='$ALICE_TOKEN'
export ALICE_USER_ID='$ALICE_USER_ID'
export BOB_EMAIL='$BOB_EMAIL'
export BOB_PASSWORD='$BOB_PASSWORD'
export BOB_TOKEN='$BOB_TOKEN'
export BOB_USER_ID='$BOB_USER_ID'
EOF

printf 'Wrote session file: %s\n' "$SCRIPT_DIR/dm_session.env"
printf 'ALICE_USER_ID=%s\n' "$ALICE_USER_ID"
printf 'BOB_USER_ID=%s\n' "$BOB_USER_ID"
