# MonyBug GitHub Issues

## 🐛 Bug Fixes

### Issue #1: Transaction Creation Fails with Column Error
**Priority**: Critical  
**Labels**: `bug`, `backend`, `database`

**Description:**
Transaction creation fails with error: `column "updated_at" of relation "transactions" does not exist`

**Steps to Reproduce:**
1. Fill out expense form in frontend
2. Submit transaction
3. Error occurs in backend

**Expected Behavior:**
Transaction should be created successfully

**Actual Behavior:**
500 error with database column error

**Solution:**
- [x] Remove `UpdatedAt` and `DeletedAt` fields from Transaction model
- [ ] Restart backend to apply changes
- [ ] Test transaction creation

---

### Issue #2: CORS Error Blocking Frontend API Calls
**Priority**: Critical  
**Labels**: `bug`, `backend`, `cors`

**Description:**
Frontend cannot access backend API due to CORS restrictions

**Steps to Reproduce:**
1. Start frontend and backend
2. Try to load categories
3. Network error in browser console

**Expected Behavior:**
API calls should work from frontend

**Actual Behavior:**
CORS error blocks all API requests

**Solution:**
- [x] Add CORS middleware to backend
- [x] Configure allowed origins for development
- [x] Test API integration

---

## ✨ Feature Requests

### Issue #3: Add Edit Transaction Functionality
**Priority**: High  
**Labels**: `enhancement`, `frontend`, `crud`

**Description:**
Users should be able to edit existing transactions

**Acceptance Criteria:**
- [ ] Add edit button to transaction items
- [ ] Create edit form modal/page
- [ ] Update transaction via API
- [ ] Refresh transaction list after update

**Technical Tasks:**
- [ ] Add `UpdateTransaction` component
- [ ] Implement PUT API call in frontend
- [ ] Add form validation for edit mode
- [ ] Handle loading states

---

### Issue #4: Implement Transaction Filtering
**Priority**: Medium  
**Labels**: `enhancement`, `frontend`, `ux`

**Description:**
Add filtering options for transactions by date range and category

**Acceptance Criteria:**
- [ ] Date range picker (from/to dates)
- [ ] Category filter dropdown
- [ ] Clear filters button
- [ ] Filter results update in real-time

**Technical Tasks:**
- [ ] Add filter UI components
- [ ] Implement client-side filtering logic
- [ ] Add URL query parameters for filters
- [ ] Persist filter state

---

### Issue #5: Add Loading States and Error Handling
**Priority**: Medium  
**Labels**: `enhancement`, `frontend`, `ux`

**Description:**
Improve user experience with loading indicators and error messages

**Acceptance Criteria:**
- [ ] Show loading spinner during API calls
- [ ] Display user-friendly error messages
- [ ] Add retry functionality for failed requests
- [ ] Show success notifications

**Technical Tasks:**
- [ ] Create Loading component
- [ ] Create ErrorMessage component
- [ ] Add loading state to all API calls
- [ ] Implement toast notifications

---

### Issue #6: Create Analytics Dashboard
**Priority**: Medium  
**Labels**: `enhancement`, `backend`, `frontend`, `analytics`

**Description:**
Add analytics page with spending insights and charts

**Acceptance Criteria:**
- [ ] Monthly spending summary
- [ ] Category breakdown chart
- [ ] Spending trends over time
- [ ] Budget vs actual comparison

**Technical Tasks:**
- [ ] Create analytics API endpoints
- [ ] Add chart library (Chart.js/Recharts)
- [ ] Create Analytics page component
- [ ] Implement data visualization

---

### Issue #7: Add Search Functionality
**Priority**: Low  
**Labels**: `enhancement`, `frontend`, `search`

**Description:**
Allow users to search transactions by description or tags

**Acceptance Criteria:**
- [ ] Search input field
- [ ] Real-time search results
- [ ] Search by description and tags
- [ ] Highlight search terms in results

**Technical Tasks:**
- [ ] Add search input component
- [ ] Implement search filtering logic
- [ ] Add search highlighting
- [ ] Optimize search performance

---

## 🚀 Backend Enhancements

### Issue #8: Create Monthly Analytics API
**Priority**: Medium  
**Labels**: `enhancement`, `backend`, `api`

**Description:**
Create API endpoints for monthly spending analytics

**Acceptance Criteria:**
- [ ] GET `/api/v1/analytics/monthly` - Monthly totals
- [ ] GET `/api/v1/analytics/categories` - Category breakdown
- [ ] GET `/api/v1/analytics/trends` - Spending trends
- [ ] Support date range parameters

**Technical Tasks:**
- [ ] Create analytics handlers
- [ ] Add database aggregation queries
- [ ] Implement caching for performance
- [ ] Add API documentation

---

### Issue #9: Implement Budget Tracking
**Priority**: Medium  
**Labels**: `enhancement`, `backend`, `budgets`

**Description:**
Add budget vs actual spending comparison

**Acceptance Criteria:**
- [ ] GET `/api/v1/budgets/{id}/status` - Budget status
- [ ] Calculate actual spending vs budget
- [ ] Return percentage used and remaining
- [ ] Support monthly/yearly budgets

**Technical Tasks:**
- [ ] Create budget status handler
- [ ] Add spending calculation logic
- [ ] Implement budget alerts
- [ ] Add budget validation

---

## 🔧 Technical Improvements

### Issue #10: Add Input Validation
**Priority**: Medium  
**Labels**: `enhancement`, `backend`, `validation`

**Description:**
Improve API input validation and error responses

**Acceptance Criteria:**
- [ ] Validate all input fields
- [ ] Return structured error responses
- [ ] Add field-specific error messages
- [ ] Prevent invalid data insertion

**Technical Tasks:**
- [ ] Add validation middleware
- [ ] Create error response structure
- [ ] Validate transaction amounts
- [ ] Validate date formats

---

### Issue #11: Setup Development Environment
**Priority**: Low  
**Labels**: `devops`, `setup`

**Description:**
Improve development setup and documentation

**Acceptance Criteria:**
- [ ] Docker compose for local development
- [ ] Environment variable documentation
- [ ] Setup scripts for new developers
- [ ] Database seeding scripts

**Technical Tasks:**
- [ ] Create Dockerfile for backend
- [ ] Create docker-compose.yml
- [ ] Add setup documentation
- [ ] Create database seed script

---

## 📱 Future Features

### Issue #12: Android Mobile App
**Priority**: Low  
**Labels**: `enhancement`, `mobile`, `android`

**Description:**
Create native Android application

**Acceptance Criteria:**
- [ ] Native Kotlin Android app
- [ ] Connect to existing backend API
- [ ] Offline transaction storage
- [ ] Sync when online

**Technical Tasks:**
- [ ] Setup Android project structure
- [ ] Implement API client
- [ ] Create UI screens
- [ ] Add offline functionality

---

## 🏷️ Labels Reference

- `bug` - Something isn't working
- `enhancement` - New feature or request
- `frontend` - Frontend/React related
- `backend` - Backend/API related
- `database` - Database related
- `mobile` - Mobile app related
- `devops` - Development operations
- `documentation` - Documentation improvements
- `critical` - Needs immediate attention
- `high` - High priority
- `medium` - Medium priority
- `low` - Low priority