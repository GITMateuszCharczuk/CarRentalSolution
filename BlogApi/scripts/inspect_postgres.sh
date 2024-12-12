#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Container name
CONTAINER="Identity_postgres"

# Database connection details
DB_NAME="Identity_db"
DB_USER="postgres"

show_help() {
    echo -e "${BLUE}PostgreSQL Database Inspector${NC}"
    echo "Usage: ./inspect_postgres.sh [option]"
    echo ""
    echo "Options:"
    echo "  1, --tables        List all tables"
    echo "  2, --users         Show user table data"
    echo "  3, --enums         List all enum types"
    echo "  4, --schema        Show table schema"
    echo "  5, --count         Show record counts for all tables"
    echo "  h, --help          Show this help message"
}

check_container() {
    if ! docker ps | grep -q $CONTAINER; then
        echo -e "${RED}Error: Container $CONTAINER is not running${NC}"
        exit 1
    fi
}

execute_query() {
    docker exec -it $CONTAINER psql -U $DB_USER -d $DB_NAME -c "$1"
}

case "$1" in
    "1"|"--tables")
        check_container
        echo -e "${GREEN}Listing all tables:${NC}"
        execute_query "\dt"
        ;;
    "2"|"--users")
        check_container
        echo -e "${GREEN}Showing users data:${NC}"
        execute_query "SELECT id, name, surname, email_address, roles, created_at FROM user_entity;"
        ;;
    "3"|"--enums")
        check_container
        echo -e "${GREEN}Listing enum types:${NC}"
        execute_query "SELECT typname, enumlabel FROM pg_enum e JOIN pg_type t ON e.enumtypid = t.oid;"
        ;;
    "4"|"--schema")
        check_container
        echo -e "${GREEN}Showing table schema:${NC}"
        execute_query "\d+ user_entity"
        ;;
    "5"|"--count")
        check_container
        echo -e "${GREEN}Showing record counts:${NC}"
        execute_query "SELECT schemaname, relname, n_live_tup FROM pg_stat_user_tables;"
        ;;
    "h"|"--help"|*)
        show_help
        ;;
esac 